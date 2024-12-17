package clients

import (
	"bytes"
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"math/rand"
	"net/http"
	"time"
)

type HttpClient struct {
	URI        string
	MaxRetries int
	HttpClient *http.Client
	limiter    *rate.Limiter
	Bearer     string
}

func (h *HttpClient) Get(path string, query map[string]string) (response *http.Response, err error) {
	var queryString = ""
	for fieldName, value := range query {
		if value == "" {
			continue
		}

		if queryString == "" {
			queryString += "?"
		} else {
			queryString += "&"
		}

		queryString += fmt.Sprintf("%s=%s", fieldName, value)
	}

	path = fmt.Sprintf("%s%s", path, queryString)

	return h.request(http.MethodGet, path, []byte(``), map[string]string{})
}

func (h *HttpClient) request(method string, path string, body []byte, headers map[string]string) (response *http.Response, err error) {
	endpoint := fmt.Sprintf("%s%s", h.URI, path)

	request, err := http.NewRequest(method, endpoint, bytes.NewReader(body))
	if err != nil {
		//@todo: add logs
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	for attribute, value := range headers {
		request.Header.Set(attribute, value)
	}

	if h.Bearer != "" {
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", h.Bearer))
	}

	return h.HttpClient.Do(request)
}

func New(reqPerSecond int) *HttpClient {
	httpTransport := http.Transport{
		TLSHandshakeTimeout: time.Duration(5) * time.Second,
	}

	httpClient := http.Client{Transport: &httpTransport, Timeout: time.Duration(10) * time.Second}

	return &HttpClient{
		HttpClient: &httpClient,
		limiter:    rate.NewLimiter(rate.Limit(reqPerSecond), 1),
	}
}

func (h *HttpClient) Do(req *http.Request) (*http.Response, error) {
	// Wait until the limiter allows another request
	err := h.limiter.Wait(context.Background())
	if err != nil {
		return nil, fmt.Errorf("rate limit error: %w", err)
	}

	for attempt := 0; attempt < h.MaxRetries; attempt++ { // Retry up to 5 times
		resp, err := h.HttpClient.Do(req)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			// Handle Retry-After
			retryAfter := resp.Header.Get("Retry-After")
			if retryAfter != "" {
				sleepTime, _ := time.ParseDuration(retryAfter + "s")
				time.Sleep(sleepTime)
				continue
			}

			// Exponential backoff
			time.Sleep(exponentialBackoff(attempt))
			continue
		}

		return resp, nil
	}

	return nil, fmt.Errorf("max retries exceeded")
}

func exponentialBackoff(attempt int) time.Duration {
	baseDelay := time.Second
	maxJitter := time.Millisecond * 100
	jitter := time.Duration(rand.Int63n(int64(maxJitter)))
	return baseDelay*(1<<attempt) + jitter
}

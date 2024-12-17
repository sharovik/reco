package clients

import "net/http"

type BaseClientInterface interface {
	Get(path string, query map[string]string) (response *http.Response, err error)
}

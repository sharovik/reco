package providers

import (
	"bytes"
	mocks "github.com/reco/mocks/pkg/clients"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestAsanaService_GetUsersListEmptyOffset(t *testing.T) {
	client := new(mocks.BaseClientInterface)

	client.On("Get", "/users", map[string]string{
		"limit": "1", "offset": "", "workspace": "",
	}).Once().Return(&http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(`{"data":[{"gid":"1208997572123832","name":"john-doe@test.lock","resource_type":"user"}],"next_page":null}`)),
	}, nil)
	s := AsanaService{Client: client}

	list, offset, err := s.GetUsersList("", 1)

	assert.Len(t, list, 1)
	assert.Equal(t, "", offset)
	assert.NoError(t, err)
}

func TestAsanaService_GetUsersListWithOffset(t *testing.T) {
	client := new(mocks.BaseClientInterface)

	client.On("Get", "/users", map[string]string{
		"limit": "1", "offset": "", "workspace": "",
	}).Once().Return(&http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(`{"data":[{"gid":"1208997305723777","name":"Pavel Simzicov","resource_type":"user"}],"next_page":{"offset":"test","path":"","uri":"https://app.asana.com/api/1.0/users?limit=3&workspace=111&offset=test"}}`)),
	}, nil)
	s := AsanaService{Client: client}

	list, offset, err := s.GetUsersList("", 1)

	assert.Len(t, list, 1)
	assert.Equal(t, "test", offset)
	assert.NoError(t, err)
}

func TestAsanaService_GetUsersListBadRequest(t *testing.T) {
	client := new(mocks.BaseClientInterface)

	var cases = map[int]string{
		http.StatusBadRequest:   "bad request received",
		http.StatusUnauthorized: "request unauthorized",
		http.StatusForbidden:    "request forbidden",
	}

	for statusCode, expectedError := range cases {
		client.On("Get", "/users", map[string]string{
			"limit": "1", "offset": "", "workspace": "",
		}).Once().Return(&http.Response{
			StatusCode: statusCode,
			Body:       io.NopCloser(bytes.NewBufferString(``)),
		}, nil)
		s := AsanaService{Client: client}

		list, offset, err := s.GetUsersList("", 1)

		assert.Len(t, list, 0)
		assert.Equal(t, "", offset)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err.Error())
	}
}

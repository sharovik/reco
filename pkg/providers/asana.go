package providers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/reco/pkg/clients"
	"github.com/reco/pkg/dto"
	"io"
	"net/http"
)

type AsanaService struct {
	Client      clients.BaseClientInterface
	WorkspaceID string
}

func (s AsanaService) GetUsersList(page string, limit int) (users []dto.UserDataItem, nextPage string, err error) {
	resp, err := s.Client.Get("/users", map[string]string{
		"workspace": s.WorkspaceID,
		"limit":     fmt.Sprintf("%d", limit),
		"offset":    page,
	})
	if err != nil {
		//@todo: add logs
		return nil, "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return nil, "", errors.New("bad request received")
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, "", errors.New("request unauthorized")
	}

	if resp.StatusCode == http.StatusForbidden {
		return nil, "", errors.New("request forbidden")
	}

	byteResp, err := io.ReadAll(resp.Body)
	if err != nil {
		//@todo: add logs
		return nil, "", err
	}

	//@todo: this can be replaced by using Pools in order to safe some memory
	response := dto.AsanaUsersListResponse{}
	if err = json.Unmarshal(byteResp, &response); err != nil {
		//@todo: add logs
		return nil, "", err
	}

	return response.Data, response.NextPage.Offset, nil
}

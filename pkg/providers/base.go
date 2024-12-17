package providers

import "github.com/reco/pkg/dto"

type BaseInterface interface {
	GetUsersList(page string, limit int) (users []dto.UserDataItem, nextPage string, err error)
}

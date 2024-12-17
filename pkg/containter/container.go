package containter

import (
	"github.com/reco/pkg/clients"
	"github.com/reco/pkg/config"
	"github.com/reco/pkg/providers"
	"github.com/reco/pkg/services"
)

// @todo: move to Pool
var c Container

type Container struct {
	IsInitialized bool
	DataProvider  providers.BaseInterface
	UsersService  services.BaseUsersServiceInterface
	Config        config.Config
}

func Get() Container {
	if c.IsInitialized {
		return c
	}

	cfg := config.New()

	c = Container{
		Config:        cfg,
		DataProvider:  loadDataProvider(cfg),
		IsInitialized: true,
		UsersService:  services.UsersService{},
	}

	return c
}

// In this case we implement only asana for now
func loadDataProvider(cfg config.Config) providers.BaseInterface {

	client := clients.New(5)

	//@todo: create a separate asana client which will extent base http
	client.URI = "https://app.asana.com/api/1.0"
	client.MaxRetries = 3
	client.TLSTimeout = 5
	client.Bearer = cfg.Asana.Token

	return providers.AsanaService{Client: client, WorkspaceID: cfg.Asana.WorkspaceID}
}

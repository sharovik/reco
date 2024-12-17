package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

const (
	envAsanaToken       = "ASANA_TOKEN"
	envAsanaWorkspaceID = "ASANA_WORKSPACE_ID"

	pathToEnv = "./.env"
)

type Config struct {
	Asana Asana
}

type Asana struct {
	Token       string
	WorkspaceID string
}

func New() Config {
	if err := godotenv.Load(pathToEnv); err != nil {
		//@todo: implement logging
		fmt.Println(err)
	}

	return Config{
		Asana: loadAsanaCredentials(),
	}
}

func loadAsanaCredentials() Asana {
	return Asana{
		Token:       os.Getenv(envAsanaToken),
		WorkspaceID: os.Getenv(envAsanaWorkspaceID),
	}
}

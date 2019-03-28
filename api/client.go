package api

import (
	"git.dtluna.net/dtluna/syncthing-cli/config"

	"gopkg.in/h2non/gentleman.v2"
)

func NewClient(cfg *config.Config) *gentleman.Client {
	return gentleman.New().
		BaseURL("http://"+cfg.Address).
		AddHeader("X-API-KEY", cfg.APIKey)
}

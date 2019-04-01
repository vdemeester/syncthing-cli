package api

import (
	"fmt"

	"git.dtluna.net/dtluna/syncthing-cli/config"

	"gopkg.in/h2non/gentleman.v2"
)

func newClient(cfg *config.Config) *gentleman.Client {
	return gentleman.New().
		BaseURL("http://"+cfg.Address).
		AddHeader("X-API-KEY", cfg.APIKey)
}

func checkResponseOK(resp *gentleman.Response) error {
	if !resp.Ok {
		return fmt.Errorf("%v %v", resp.StatusCode, resp.String())
	}
	return nil
}

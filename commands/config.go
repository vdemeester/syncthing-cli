package commands

import (
	"fmt"

	"git.dtluna.net/dtluna/syncthing-cli/api"
	"git.dtluna.net/dtluna/syncthing-cli/config"
)

func Config(cfg *config.Config) error {
	stconfig, err := api.GetConfig(cfg)
	if err != nil {
		return err
	}

	fmt.Println(stconfig)
	return nil
}

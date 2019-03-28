package commands

import (
	"fmt"

	"git.dtluna.net/dtluna/syncthing-cli/api"
	"git.dtluna.net/dtluna/syncthing-cli/config"

	"github.com/hashicorp/errwrap"
)

func Config(cfg *config.Config) error {
	stconfig, err := api.GetConfig(cfg)
	if err != nil {
		return errwrap.Wrapf("getting config: {{err}}", err)
	}

	fmt.Println(stconfig)
	return nil
}
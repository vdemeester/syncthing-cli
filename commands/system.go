package commands

import (
	"git.dtluna.net/dtluna/syncthing-cli/api"
	"git.dtluna.net/dtluna/syncthing-cli/config"

	"github.com/hashicorp/errwrap"
)

func Restart(cfg *config.Config) error {
	err := api.Restart(cfg)
	if err != nil {
		return errwrap.Wrapf("requesting a restart: {{err}}", err)
	}
	return nil
}

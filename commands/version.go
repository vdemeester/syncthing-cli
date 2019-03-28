package commands

import (
	"fmt"

	"git.dtluna.net/dtluna/syncthing-cli/api"
	"git.dtluna.net/dtluna/syncthing-cli/config"

	"github.com/hashicorp/errwrap"
)

func Version(cfg *config.Config) error {
	info, err := api.Version(cfg)
	if err != nil {
		return errwrap.Wrapf("requesting version: {{err}}", err)

	}
	fmt.Println(info.LongVersion)

	return nil
}

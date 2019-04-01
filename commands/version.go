package commands

import (
	"fmt"

	"git.dtluna.net/dtluna/syncthing-cli/api"
	"git.dtluna.net/dtluna/syncthing-cli/config"
)

func Version(cfg *config.Config) error {
	info, err := api.Version(cfg)
	if err != nil {
		return err
	}
	fmt.Println(info.LongVersion)

	return nil
}

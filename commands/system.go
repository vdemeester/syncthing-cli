package commands

import (
	"git.dtluna.net/dtluna/syncthing-cli/api"
	"git.dtluna.net/dtluna/syncthing-cli/config"
)

func Restart(cfg *config.Config) error {
	return api.Restart(cfg)
}

func Shutdown(cfg *config.Config) error {
	return api.Shutdown(cfg)
}

func Pause(cfg *config.Config, devices []string) error {
	if len(devices) == 0 {
		return api.Pause(cfg, "")
	}

	for _, device := range devices {
		err := api.Pause(cfg, device)
		if err != nil {
			return err
		}
	}
	return nil
}

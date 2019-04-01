package api

import (
	"net/http"

	"git.dtluna.net/dtluna/syncthing-cli/config"
)

const (
	RestartPath  = "/rest/system/restart"
	ShutdownPath = "/rest/system/shutdown"
)

func Restart(cfg *config.Config) error {
	req := NewClient(cfg).Request().Path(RestartPath).Method(http.MethodPost)
	resp, err := req.Send()
	if err != nil {
		return err
	}

	err = checkResponseOK(resp)
	if err != nil {
		return err
	}

	return nil
}

func Shutdown(cfg *config.Config) error {
	req := NewClient(cfg).Request().Path(ShutdownPath).Method(http.MethodPost)
	resp, err := req.Send()
	if err != nil {
		return err
	}

	err = checkResponseOK(resp)
	if err != nil {
		return err
	}

	return nil
}

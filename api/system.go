package api

import (
	"net/http"

	"git.dtluna.net/dtluna/syncthing-cli/config"
)

const (
	restartPath  = "/rest/system/restart"
	shutdownPath = "/rest/system/shutdown"
	pausePath    = "/rest/system/pause"
	resumePath   = "/rest/system/resume"

	restartError  = "requesting a restart: {{err}}"
	pauseError    = "requesting to pause a device: {{err}}"
	shutdownError = "requesting a shutdown: {{err}}"
	resumeError   = "requesting to resume a device: {{err}}"
)

func Restart(cfg *config.Config) error {
	req := newClient(cfg).Request().Path(restartPath).Method(http.MethodPost)
	resp, err := req.Send()
	if err != nil {
		return wrapError(err, restartError)
	}

	err = checkResponseOK(resp)
	if err != nil {
		return wrapError(err, restartError)
	}

	return nil
}

func Shutdown(cfg *config.Config) error {
	req := newClient(cfg).Request().Path(shutdownPath).Method(http.MethodPost)
	resp, err := req.Send()
	if err != nil {
		return wrapError(err, shutdownError)
	}

	err = checkResponseOK(resp)
	if err != nil {
		return wrapError(err, shutdownError)
	}

	return nil
}

func Pause(cfg *config.Config, device string) error {
	req := newClient(cfg).Request().Path(pausePath).Method(http.MethodPost)
	if device != "" {
		req = req.AddQuery("device", device)
	}

	resp, err := req.Send()
	if err != nil {
		return wrapError(err, pauseError)
	}

	err = checkResponseOK(resp)
	if err != nil {
		return wrapError(err, pauseError)
	}

	return nil
}

func Resume(cfg *config.Config, device string) error {
	req := newClient(cfg).Request().Path(resumePath).Method(http.MethodPost)
	if device != "" {
		req = req.AddQuery("device", device)
	}

	resp, err := req.Send()
	if err != nil {
		return wrapError(err, resumeError)
	}

	err = checkResponseOK(resp)
	if err != nil {
		return wrapError(err, resumeError)
	}

	return nil
}

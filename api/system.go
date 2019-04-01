package api

import (
	"net/http"

	"git.dtluna.net/dtluna/syncthing-cli/config"
)

const (
	RestartPath  = "/rest/system/restart"
	ShutdownPath = "/rest/system/shutdown"
	PausePath    = "/rest/system/pause"
	ResumePath   = "/rest/system/resume"

	RestartError  = "requesting a restart: {{err}}"
	PauseError    = "requesting to pause a device: {{err}}"
	ShutdownError = "requesting a shutdown: {{err}}"
	ResumeError   = "requesting to resume a device: {{err}}"
)

func Restart(cfg *config.Config) error {
	req := NewClient(cfg).Request().Path(RestartPath).Method(http.MethodPost)
	resp, err := req.Send()
	if err != nil {
		return wrapError(err, RestartError)
	}

	err = checkResponseOK(resp)
	if err != nil {
		return wrapError(err, RestartError)
	}

	return nil
}

func Shutdown(cfg *config.Config) error {
	req := NewClient(cfg).Request().Path(ShutdownPath).Method(http.MethodPost)
	resp, err := req.Send()
	if err != nil {
		return wrapError(err, ShutdownError)
	}

	err = checkResponseOK(resp)
	if err != nil {
		return wrapError(err, ShutdownError)
	}

	return nil
}

func Pause(cfg *config.Config, device string) error {
	req := NewClient(cfg).Request().Path(PausePath).Method(http.MethodPost)
	if device != "" {
		req = req.AddQuery("device", device)
	}

	resp, err := req.Send()
	if err != nil {
		return wrapError(err, PauseError)
	}

	err = checkResponseOK(resp)
	if err != nil {
		return wrapError(err, PauseError)
	}

	return nil
}

func Resume(cfg *config.Config, device string) error {
	req := NewClient(cfg).Request().Path(ResumePath).Method(http.MethodPost)
	if device != "" {
		req = req.AddQuery("device", device)
	}

	resp, err := req.Send()
	if err != nil {
		return wrapError(err, ResumeError)
	}

	err = checkResponseOK(resp)
	if err != nil {
		return wrapError(err, ResumeError)
	}

	return nil
}

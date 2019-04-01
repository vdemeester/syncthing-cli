package api

import (
	"git.dtluna.net/dtluna/syncthing-cli/config"
)

const (
	versionPath = "/rest/system/version"

	getVersionError = "requesting version: {{err}}"
)

type VersionInfo struct {
	Architecture string `json:"arch"`
	Codename     string
	IsBeta       bool
	IsCandidate  bool
	IsRelease    bool
	LongVersion  string
	OS           string
	Version      string
}

func Version(cfg *config.Config) (*VersionInfo, error) {
	req := newClient(cfg).Request().Path(versionPath)
	resp, err := req.Send()
	if err != nil {
		return nil, wrapError(err, getVersionError)
	}

	err = checkResponseOK(resp)
	if err != nil {
		return nil, wrapError(err, getVersionError)
	}

	info := new(VersionInfo)
	err = resp.JSON(info)
	if err != nil {
		return nil, wrapError(err, getVersionError)
	}

	return info, nil
}

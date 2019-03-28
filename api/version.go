package api

import (
	"git.dtluna.net/dtluna/syncthing-cli/config"
)

const (
	VersionPath = "/rest/system/version"
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
	req := NewClient(cfg).Request().Path(VersionPath)
	resp, err := req.Send()
	if err != nil {
		return nil, err
	}

	err = checkResponseOK(resp)
	if err != nil {
		return nil, err
	}

	info := new(VersionInfo)
	err = resp.JSON(info)
	if err != nil {
		return nil, err
	}

	return info, nil
}

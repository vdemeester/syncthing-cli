package api

import (
	"fmt"

	"git.dtluna.net/dtluna/syncthing-cli/config"
)

const (
	deviceStatsPath = "/rest/stats/device"
	folderStatsPath = "/rest/stats/folder"

	getDeviceStatsError = "getting device stats: {{err}}"
	getFolderStatsError = "getting folder stats: {{err}}"
)

type DeviceStats struct {
	LastSeen string
}

func GetDeviceStats(cfg *config.Config) (*map[string]DeviceStats, error) {
	req := newClient(cfg).Request().Path(deviceStatsPath)
	resp, err := req.Send()
	if err != nil {
		return nil, wrapError(err, getDeviceStatsError)
	}

	err = checkResponseOK(resp)
	if err != nil {
		return nil, wrapError(err, getDeviceStatsError)
	}

	stats := new(map[string]DeviceStats)
	err = resp.JSON(stats)
	if err != nil {
		return nil, wrapError(err, getDeviceStatsError)
	}

	return stats, nil
}

type FolderStatsFile struct {
	Filename string
	At       string
}

func (fsf FolderStatsFile) String() string {
	return fmt.Sprintf(
		`Filename: %q
At: %v`,
		fsf.Filename,
		fsf.At,
	)
}

type FolderStats struct {
	LastScan string
	LastFile FolderStatsFile
}

func (fs FolderStats) String() string {
	return fmt.Sprintf(
		`Last scan: %v
Last file: %v`,
		fs.LastScan,
		Indent(fs.LastFile.String(), 2),
	)
}

func GetFolderStats(cfg *config.Config) (*map[string]FolderStats, error) {
	req := newClient(cfg).Request().Path(folderStatsPath)
	resp, err := req.Send()
	if err != nil {
		return nil, wrapError(err, getFolderStatsError)
	}

	err = checkResponseOK(resp)
	if err != nil {
		return nil, wrapError(err, getFolderStatsError)
	}

	stats := new(map[string]FolderStats)
	err = resp.JSON(stats)
	if err != nil {
		return nil, wrapError(err, getFolderStatsError)
	}

	return stats, nil
}

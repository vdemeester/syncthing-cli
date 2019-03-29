package api

import (
	"fmt"

	"git.dtluna.net/dtluna/syncthing-cli/config"
)

const (
	DeviceStatsPath = "/rest/stats/device"
	FolderStatsPath = "/rest/stats/folder"
)

type DeviceStats struct {
	LastSeen string
}

func GetDeviceStats(cfg *config.Config) (*map[string]DeviceStats, error) {
	req := NewClient(cfg).Request().Path(DeviceStatsPath)
	resp, err := req.Send()
	if err != nil {
		return nil, err
	}

	err = checkResponseOK(resp)
	if err != nil {
		return nil, err
	}

	stats := new(map[string]DeviceStats)
	err = resp.JSON(stats)
	if err != nil {
		return nil, err
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
	req := NewClient(cfg).Request().Path(FolderStatsPath)
	resp, err := req.Send()
	if err != nil {
		return nil, err
	}

	err = checkResponseOK(resp)
	if err != nil {
		return nil, err
	}

	stats := new(map[string]FolderStats)
	err = resp.JSON(stats)
	if err != nil {
		return nil, err
	}

	return stats, nil
}

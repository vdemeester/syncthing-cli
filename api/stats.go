package api

import (
	"git.dtluna.net/dtluna/syncthing-cli/config"
)

const DeviceStatsPath = "/rest/stats/device"

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

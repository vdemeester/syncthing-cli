package commands

import (
	"fmt"

	"git.dtluna.net/dtluna/syncthing-cli/api"
	"git.dtluna.net/dtluna/syncthing-cli/config"
	"git.dtluna.net/dtluna/syncthing-cli/format"
)

func DeviceList(cfg *config.Config) error {
	stconfig, err := api.GetConfig(cfg)
	if err != nil {
		return err
	}

	fmt.Println(format.IndentDevices(stconfig.Devices, 2))
	return nil
}

func DeviceStats(cfg *config.Config) error {
	stats, err := api.GetDeviceStats(cfg)
	if err != nil {
		return err
	}
	for name, devStats := range *stats {
		fmt.Printf("%v:\n  Last seen: %v\n", name, devStats.LastSeen)
	}
	return nil
}

func DeviceAdd(
	cfg *config.Config,
	deviceID, deviceName, compression, certName string,
	addresses []string,
	introducer bool,
) error {
	stconfig, err := api.GetConfig(cfg)
	if err != nil {
		return err
	}

	newDevice := api.Device{
		DeviceID:    deviceID,
		Name:        deviceName,
		Addresses:   addresses,
		Compression: compression,
		CertName:    certName,
		Introducer:  introducer,
	}

	stconfig.Devices = append(stconfig.Devices, newDevice)

	stconfig, err = api.SetConfig(cfg, stconfig)
	if err != nil {
		return err
	}
	fmt.Println(stconfig)

	return nil
}

package commands

import (
	"fmt"

	"git.dtluna.net/dtluna/syncthing-cli/api"
	"git.dtluna.net/dtluna/syncthing-cli/config"
	"git.dtluna.net/dtluna/syncthing-cli/format"
)

const (
	deviceNotFoundErrorTemplate = "no device with ID %q"
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

	_, err = api.SetConfig(cfg, stconfig)
	if err != nil {
		return err
	}

	return nil
}

func removeDevice(devices []api.Device, deviceID string) ([]api.Device, error) {
	for i, device := range devices {
		if device.DeviceID == deviceID {
			return append(devices[:i], devices[i+1:]...), nil
		}
	}
	return nil, fmt.Errorf(deviceNotFoundErrorTemplate, deviceID)
}

func DeviceRemove(cfg *config.Config, deviceID string) error {
	stconfig, err := api.GetConfig(cfg)
	if err != nil {
		return err
	}

	stconfig.Devices, err = removeDevice(stconfig.Devices, deviceID)
	if err != nil {
		return err
	}

	_, err = api.SetConfig(cfg, stconfig)
	if err != nil {
		return err
	}

	return nil
}

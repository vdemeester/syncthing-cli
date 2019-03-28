package config

import (
	"net"

	"git.dtluna.net/dtluna/syncthing-cli/constants"
	"gopkg.in/ini.v1"
)

type Config struct {
	APIKey  string `ini:"api-key"`
	Address string `ini:"address"`
}

func CreateBlankConfigFile(path string) error {
	file := ini.Empty()
	section := file.Section("")

	key, err := section.NewKey("api-key", "")
	if err != nil {
		return err
	}
	key.Comment = "Specify the API key below"

	_, err = section.NewKey("address", constants.DefaultAddress)
	if err != nil {
		return err
	}

	return file.SaveTo(path)
}

func Parse(path string) (*Config, error) {
	file, err := ini.Load(path)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	err = file.MapTo(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func Merge(old Config, address *net.TCPAddr, APIKey string) (*Config, error) {
	if APIKey != "" {
		old.APIKey = APIKey
	}
	if address != nil {
		old.Address = address.String()
	}

	//validate the specified address
	_, err := net.ResolveTCPAddr("tcp", old.Address)
	if err != nil {
		return nil, err
	}

	return &old, nil
}

package main

import (
	"fmt"
	"os"
	"path"

	"git.dtluna.net/dtluna/syncthing-cli/commands"
	"git.dtluna.net/dtluna/syncthing-cli/config"
	"git.dtluna.net/dtluna/syncthing-cli/constants"

	"github.com/OpenPeeDeeP/xdg"
	"github.com/Unknwon/com"
	"gopkg.in/alecthomas/kingpin.v2"
)

func printToStderr(v interface{}) {
	fmt.Fprintf(
		os.Stderr,
		"%v: error: creating default config file: %v\n",
		constants.AppName,
		v,
	)
}

func fatalIfError(err error) {
	if err != nil {
		printToStderr(err)
		os.Exit(1)
	}
}

func main() {
	x := xdg.New("", constants.AppName)
	configDir := x.ConfigHome()
	defaultConfigPath := path.Join(configDir, "config.ini")
	if !com.IsExist(defaultConfigPath) {
		if !com.IsExist(configDir) {
			err := os.MkdirAll(configDir, os.ModePerm)
			fatalIfError(err)
		}
		err := config.CreateBlankConfigFile(defaultConfigPath)
		fatalIfError(err)
	}

	app := kingpin.New(constants.AppName, "CLI client for Syncthing")
	app.Version(constants.Version)

	configPath := app.Flag("config", "Location of the config file.").
		Short('c').
		Default(defaultConfigPath).
		ExistingFile()

	address := app.Flag("address", "Address of the Syncthing daemon.").
		Short('a').
		TCP()

	apiKey := app.Flag("api-key", "API key to access the REST API of the Syncthing daemon.").
		Short('k').
		String()

	version := app.Command("version", "Show the current Syncthing version information.").Alias("v")

	configCmd := app.Command("config", "Show the current configuration.").Alias("c").Alias("conf")

	device := app.Command("device", "Work with devices.").Alias("d").Alias("dev")
	deviceList := device.Command("list", "List devices.").Alias("l").Alias("ls")
	deviceStats := device.Command("stats", "Show device stats.").Alias("s").Alias("st")

	folder := app.Command("folder", "Work with folders.").Alias("f").Alias("fl").Alias("fold")
	folderList := folder.Command("list", "List folders.").Alias("l").Alias("ls")
	folderStats := folder.Command("stats", "Show folder stats.").Alias("s").Alias("st")

	commandName := kingpin.MustParse(app.Parse(os.Args[1:]))

	cfg, err := config.Parse(*configPath)
	app.FatalIfError(err, "parsing config")

	cfg, err = config.Merge(*cfg, *address, *apiKey)
	app.FatalIfError(err, "merging config with arguments")

	err = func() error {
		switch commandName {
		case configCmd.FullCommand():
			return commands.Config(cfg)
		case version.FullCommand():
			return commands.Version(cfg)
		case deviceList.FullCommand():
			return commands.DeviceList(cfg)
		case deviceStats.FullCommand():
			return commands.DeviceStats(cfg)
		case folderList.FullCommand():
			return commands.FolderList(cfg)
		case folderStats.FullCommand():
			return commands.FolderStats(cfg)
		}
		return nil
	}()
	app.FatalIfError(err, "")
}

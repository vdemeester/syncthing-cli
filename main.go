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

	deviceAdd := device.Command("add", "Add a new device.").Alias("a")
	deviceAddID := deviceAdd.Arg("ID", "ID of the new device.").Required().String()
	deviceAddName := deviceAdd.Arg("name", "Name of the new device.").String()
	deviceAddAddresses := deviceAdd.Flag("addresses", "Addresses of the new device.").
		Default(constants.DynamicAddress).
		Strings()
	deviceAddIntroducer := deviceAdd.Flag("introducer", "Mark device as an introducer.").
		Short('i').
		Bool()
	deviceAddCompression := deviceAdd.Flag("compression", "Specify the compression to use.").
		Default(constants.CompressionMetadataOnly).
		Enum(
			constants.CompressionMetadataOnly,
			constants.CompressionAllData,
			constants.CompressionOff,
		)
	deviceAddCertName := deviceAdd.Flag("cert-name", "Specify the certificate name.").String()

	deviceRemove := device.Command("remove", "Remove a device.").Alias("r")
	deviceRemoveID := deviceRemove.Arg("ID", "ID of the device to remove.").Required().String()

	folder := app.Command("folder", "Work with folders.").Alias("f").Alias("fl").Alias("fold")
	folderList := folder.Command("list", "List folders.").Alias("l").Alias("ls")
	folderStats := folder.Command("stats", "Show folder stats.").Alias("s").Alias("st")

	restart := app.Command("restart", "Restart the Syncthing daemon.")
	shutdown := app.Command("shutdown", "Shutdown the Syncthing daemon.")

	pause := app.Command("pause", "Pause the given devices or all devices.").Alias("p")
	pauseDevices := pause.Arg("devices", "Devices to pause. If non specified all devices are paused.").
		Strings()

	resume := app.Command("resume", "Resume the given devices or all devices.").Alias("r")
	resumeDevices := resume.Arg("devices", "Devices to resume. If non specified all devices are resumed.").
		Strings()

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
		case deviceAdd.FullCommand():
			return commands.DeviceAdd(
				cfg,
				*deviceAddID,
				*deviceAddName,
				*deviceAddCompression,
				*deviceAddCertName,
				*deviceAddAddresses,
				*deviceAddIntroducer,
			)
		case deviceRemove.FullCommand():
			return commands.DeviceRemove(cfg, *deviceRemoveID)
		case folderList.FullCommand():
			return commands.FolderList(cfg)
		case folderStats.FullCommand():
			return commands.FolderStats(cfg)
		case restart.FullCommand():
			return commands.Restart(cfg)
		case shutdown.FullCommand():
			return commands.Shutdown(cfg)
		case pause.FullCommand():
			return commands.Pause(cfg, *pauseDevices)
		case resume.FullCommand():
			return commands.Resume(cfg, *resumeDevices)
		}
		return nil
	}()
	app.FatalIfError(err, "")
}

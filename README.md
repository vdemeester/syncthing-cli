# syncthing-cli

CLI client for Syncthing in Go

## Installation

To install or update the gorgeous binary into your `$GOPATH` as usual, run:
```shell
go get -u git.dtluna.net/dtluna/syncthing-cli
```

Or you can download the binaries from [the releases page](https://git.dtluna.net/dtluna/syncthing-cli/releases)


```shell
$ syncthing-cli 
usage: syncthing-cli [<flags>] <command> [<args> ...]

CLI client for Syncthing

Flags:
      --help             Show context-sensitive help (also try --help-long and --help-man).
      --version          Show application version.
  -c, --config=/home/dt/.config/syncthing-cli/config.ini  
                         Location of the config file.
  -a, --address=ADDRESS  Address of the Syncthing daemon.
  -k, --api-key=API-KEY  API key to access the REST API of the Syncthing daemon.

Commands:
  help [<command>...]
    Show help.

  version
    Show the current Syncthing version information.

  config
    Show the current configuration.

  device list
    List devices.

  device stats
    Show device stats.

  device add [<flags>] <ID> [<name>]
    Add a new device.

  device remove <ID>
    Remove a device.

  folder list
    List folders.

  folder stats
    Show folder stats.

  folder add [<flags>] <label> <path>
    Add a new folder.

  folder remove <ID>
    Remove a folder.

  restart
    Restart the Syncthing daemon.

  shutdown
    Shutdown the Syncthing daemon.

  pause [<devices>...]
    Pause the given devices or all devices.

  resume [<devices>...]
    Resume the given devices or all devices.

```

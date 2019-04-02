package commands

import (
	"fmt"
	"path"

	"git.dtluna.net/dtluna/syncthing-cli/api"
	"git.dtluna.net/dtluna/syncthing-cli/config"
	"git.dtluna.net/dtluna/syncthing-cli/format"
)

const (
	folderNotFoundErrorTemplate = "no folder with ID %q"
)

func FolderList(cfg *config.Config) error {
	stconfig, err := api.GetConfig(cfg)
	if err != nil {
		return err
	}

	fmt.Println(format.IndentFolders(stconfig.Folders, 2))
	return nil
}

func FolderStats(cfg *config.Config) error {
	stats, err := api.GetFolderStats(cfg)
	if err != nil {
		return err
	}

	for name, folderStats := range *stats {
		fmt.Printf("%v: %v\n", name, api.Indent(folderStats.String(), 2))
	}
	return nil
}

func removeFolder(folders []api.Folder, folderID string) ([]api.Folder, error) {
	for i, folder := range folders {
		if folder.ID == folderID {
			return append(folders[:i], folders[i+1:]...), nil
		}
	}
	return nil, fmt.Errorf(folderNotFoundErrorTemplate, folderID)
}

func FolderRemove(cfg *config.Config, folderID string) error {
	stconfig, err := api.GetConfig(cfg)
	if err != nil {
		return err
	}

	stconfig.Folders, err = removeFolder(stconfig.Folders, folderID)
	if err != nil {
		return err
	}

	_, err = api.SetConfig(cfg, stconfig)
	if err != nil {
		return err
	}

	return nil
}

func FolderAdd(
	cfg *config.Config,
	label, ID, folderPath, folderType, order string,
	shareWith []string,
	minDiskFreePct, rescanIntervalS int,
	fsWatcherEnabled, ignorePerms, ignoreDelete bool,
) error {
	if ID == "" {
		ID = path.Base(folderPath)
	}

	folderDevices := []api.FolderDevice{}
	for _, deviceID := range shareWith {
		folderDevices = append(folderDevices, api.FolderDevice{DeviceID: deviceID})
	}

	newFolder := api.Folder{
		ID:               ID,
		Label:            label,
		Path:             folderPath,
		Type:             folderType,
		Order:            order,
		IgnorePerms:      ignorePerms,
		Devices:          folderDevices,
		MinDiskFreePct:   minDiskFreePct,
		RescanIntervalS:  rescanIntervalS,
		FSWatcherEnabled: fsWatcherEnabled,
		IgnoreDelete:     ignoreDelete,
	}
	fmt.Println(newFolder)
	return nil
}

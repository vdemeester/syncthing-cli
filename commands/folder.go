package commands

import (
	"fmt"

	"git.dtluna.net/dtluna/syncthing-cli/api"
	"git.dtluna.net/dtluna/syncthing-cli/config"
	"git.dtluna.net/dtluna/syncthing-cli/format"

	"github.com/hashicorp/errwrap"
)

func FolderList(cfg *config.Config) error {
	stconfig, err := api.GetConfig(cfg)
	if err != nil {
		return errwrap.Wrapf("getting config: {{err}}", err)
	}

	fmt.Println(format.IndentFolders(stconfig.Folders, 2))
	return nil
}

func FolderStats(cfg *config.Config) error {
	stats, err := api.GetFolderStats(cfg)
	if err != nil {
		return errwrap.Wrapf("getting folder stats: {{err}}", err)
	}

	for name, folderStats := range *stats {
		fmt.Printf("%v: %v\n", name, api.Indent(folderStats.String(), 2))
	}
	return nil
}

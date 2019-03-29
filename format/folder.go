package format

import (
	"fmt"
	"strings"

	"git.dtluna.net/dtluna/syncthing-cli/api"
)

func IndentFolders(ds []api.Folder, depth int) string {
	if len(ds) == 0 {
		return "None"
	}

	sep := strings.Repeat(" ", depth)
	lines := []string{}

	for i, el := range ds {
		line := strings.ReplaceAll(fmt.Sprintf("%v. %s", i+1, el), "\n", "\n"+sep)
		lines = append(lines, line)
	}
	return "\n" + strings.Join(lines, "\n")
}

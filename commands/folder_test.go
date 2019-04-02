package commands

import (
	"fmt"
	"testing"

	"git.dtluna.net/dtluna/syncthing-cli/api"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RemoveFolderTestSuite struct {
	suite.Suite
}

func (suite *RemoveFolderTestSuite) TestError() {
	t := suite.T()
	folders := []api.Folder{}
	folderID := "some-id"
	result, err := removeFolder(folders, folderID)
	assert.Nil(t, result)
	assert.EqualError(t, err, fmt.Sprintf(folderNotFoundErrorTemplate, folderID))
}

func (suite *RemoveFolderTestSuite) TestSuccess() {
	t := suite.T()
	folderID := "some-id"
	folders := []api.Folder{
		api.Folder{
			ID: folderID,
		},
	}
	result, err := removeFolder(folders, folderID)
	assert.NoError(t, err)
	assert.Empty(t, result)
}

func TestRemoveFolderTestSuite(t *testing.T) {
	suite.Run(t, new(RemoveFolderTestSuite))
}

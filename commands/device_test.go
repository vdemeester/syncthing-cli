package commands

import (
	"fmt"
	"testing"

	"git.dtluna.net/dtluna/syncthing-cli/api"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RemoveDeviceTestSuite struct {
	suite.Suite
}

func (suite *RemoveDeviceTestSuite) TestError() {
	t := suite.T()
	devices := []api.Device{}
	deviceID := "some-id"
	result, err := removeDevice(devices, deviceID)
	assert.Nil(t, result)
	assert.EqualError(t, err, fmt.Sprintf(deviceNotFoundErrorTemplate, deviceID))
}

func (suite *RemoveDeviceTestSuite) TestSuccess() {
	t := suite.T()
	deviceID := "some-id"
	devices := []api.Device{
		api.Device{
			DeviceID: deviceID,
		},
	}
	result, err := removeDevice(devices, deviceID)
	assert.NoError(t, err)
	assert.Empty(t, result)
}

func TestRemoveDeviceTestSuite(t *testing.T) {
	suite.Run(t, new(RemoveDeviceTestSuite))
}

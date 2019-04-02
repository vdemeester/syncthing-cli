package api

import (
	"git.dtluna.net/dtluna/syncthing-cli/config"
)

const (
	randomStringPath = "/rest/svc/random/string"

	getRandomStringError = "requesting a random string: {{err}}"
)

type randomStringResult struct {
	Random string
}

func RandomString(cfg *config.Config, length uint) (string, error) {
	req := newClient(cfg).Request().Path(randomStringPath)
	resp, err := req.Send()
	if err != nil {
		return "", wrapError(err, getRandomStringError)
	}

	err = checkResponseOK(resp)
	if err != nil {
		return "", wrapError(err, getRandomStringError)
	}

	result := new(randomStringResult)
	err = resp.JSON(result)
	if err != nil {
		return "", wrapError(err, getRandomStringError)
	}

	return result.Random, nil
}

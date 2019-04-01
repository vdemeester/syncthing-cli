package api

import (
	"github.com/hashicorp/errwrap"
)

func wrapError(err error, message string) error {
	return errwrap.Wrapf(message, err)
}

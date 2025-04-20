package server

import (
	"errors"
)

func PublicWrapError(err error, prefix string) error {
	return errors.Join(errors.New(prefix), err)
}

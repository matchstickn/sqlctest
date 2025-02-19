package server

import (
	"errors"
	// "github.com/gofiber/fiber/v2"
)

func PublicWrapError(err error, prefix string) error {
	return errors.Join(errors.New(prefix), err)
}

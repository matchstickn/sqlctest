package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func CheckAuthentication(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return err
		}

		if sess.Get("state") == nil {
			return c.Redirect("/")
		} else {
			return c.Next()
		}

	}
}

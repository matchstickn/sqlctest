package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// Login logic
func StartAuthenticatedSession(auth *Authenticator, store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return err
		}

		state, err := generateRandomState()
		if err != nil {
			return err
		}

		sess.Set("sessionState", state)
		sess.SetExpiry(time.Minute * 2)
		if err := sess.Save(); err != nil {
			return err
		}

		return c.Next()
	}
}

// Logout logic
func StopAuthenticatedSession(auth *Authenticator, store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {

	}
}

func AuthenticationCallback(auth *Authenticator, store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return err
		}
		if c.Query("sessionState") != sess.Get("sessionState") {
			return err
		}

		token, err := auth.Exchange(context.Background(), "")
		if err != nil {
			return err
		}

		idToken, err := auth.VerifyIDToken(context.Background(), token)
		if err != nil {
			return err
		}

		var profile map[string]interface{}
		if err := idToken.Claims(&profile); err != nil {
			return err
		}

		sess.Set("accessToken", token.AccessToken)
		sess.Set("profile", profile)
		if err := sess.Save(); err != nil {
			return err
		}

		return c.Redirect("/")
	}
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}

package auth

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/matchstickn/sqlctest/internal/auth/fiber_goth"
)

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

func GothSetUpRoutes(app *fiber.App) error {
	fiber_goth.SessionStore = session.New(session.Config{
		Expiration: time.Hour * 12,
	})

	if err := godotenv.Load(); err != nil {
		return err
	}

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), "https://127.0.0.1:4000/list"),
	)

	app.Get("/login/:provider", GothAuthenticate)
	app.Get("/auth/callback/:provider", GothAuthenitcationCallback)
	app.Get("/logout", GothLogout)

	return nil
}

func GothAuthenticate(c *fiber.Ctx) error {
	user, err := fiber_goth.CompleteUserAuth(c, fiber_goth.CompleteUserAuthOptions{
		ShouldLogout: true,
	})
	if err != nil {
		return fiber_goth.BeginAuthHandler(c)
	}
	return c.JSON(user)

}

func GothAuthenitcationCallback(c *fiber.Ctx) error {
	user, err := fiber_goth.CompleteUserAuth(c, fiber_goth.CompleteUserAuthOptions{
		ShouldLogout: true,
	})
	if err != nil {
		return c.SendString(err.Error())
	}

	fmt.Println(user)

	return c.JSON(user)
}

func GothLogout(c *fiber.Ctx) error {
	fiber_goth.Logout(c)
	return c.Redirect("/list", fiber.StatusFound)
}

func GothGetProviderIndex() (*ProviderIndex, error) {
	services := map[string]string{
		"google": "Google",
	}

	var keys []string
	for key := range services {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	providerIndex := &ProviderIndex{
		Providers:    keys,
		ProvidersMap: services,
	}

	return providerIndex, nil
}

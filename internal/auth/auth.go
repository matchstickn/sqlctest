package auth

import (
	"os"
	"sort"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/matchstickn/sqlctest/internal/auth/fiber_goth"
)

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

func GothGetProviderIndex() (*ProviderIndex, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), "http:localhost:3000/auth/google/callback"),
	)

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
	return c.JSON(user)
}

func GothLogout(c *fiber.Ctx) error {
	fiber_goth.Logout(c)
	return c.Redirect("/", fiber.StatusFound)
}

package auth

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Login(c *fiber.Ctx) error {
	user := c.FormValue("user")

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  user,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	secret, ok := os.LookupEnv("AUTH0_CLIENT_SECRET")
	if !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func Restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}

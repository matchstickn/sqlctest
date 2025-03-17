package auth

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gofiber/fiber/v2"
	"github.com/matchstickn/sqlctest/assets/db"
)

type CustomClaims struct {
	UserID string `json:"sub"`
	Name   string `json:"name"`
	Admin  bool   `json:"admin,omitempty"`
}

func New(query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")

		spinner, err := retriveUserFromToken(query, token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   "unauthorized",
			})
		}

		return c.JSON(spinner)
	}
}

func validateToken(token string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	// Set default values for missing claims
	if claims.Admin == false { // Default to non-admin if claim missing
		claims.Admin = false
	}

	return claims, nil
}

func retriveUserFromToken(query *db.Queries, token string) (db.Spinner, error) {
	spinner, err := query.RetriveSpinner(context.Background(), token)
	if err != nil {
		return db.Spinner{}, err
	}
	return spinner, nil
}

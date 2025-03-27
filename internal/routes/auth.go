package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/matchstickn/sqlctest/internal/auth"
)

func Test(c *fiber.Ctx) error {
	return c.JSON("testing")
}

func SignUp() fiber.Handler {
	return func(c *fiber.Ctx) error {
		signUpReq := auth.SignUpRequest{}
		if err := c.BodyParser(&signUpReq); err != nil {
			return err
		}

		id, validated, err := signUpReq.GetSignUpId()
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"id":              id,
			"email_validated": validated,
		})
	}
}

func GetAuthenticatedAccessToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessTokenReq := auth.AccessTokenRequest{}
		if err := c.BodyParser(&accessTokenReq); err != nil {
			return err
		}

		accessToken, err := accessTokenReq.GetAccessToken()
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"access_token": accessToken,
		})
	}
}

// Uses AccessTokenReponse struct
func SocialLoginCallback() fiber.Handler {
	return func(c *fiber.Ctx) error {
		qu := c.Queries()

		accessTokenRespMarshel, err := json.Marshal(&qu)
		if err != nil {
			return err
		}

		accessTokenResp := auth.AccessTokenResponse{}
		if err := json.Unmarshal(accessTokenRespMarshel, &accessTokenResp); err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"access_token": accessTokenResp.AccessToken,
		})
	}
}

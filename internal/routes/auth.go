package routes

import (
	"fmt"
	"os"

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

		// accessTokenRespMarshel, err := json.Marshal(&qu)
		// if err != nil {
		// 	return err
		// }

		// accessTokenResp := auth.AccessTokenResponse{}
		// if err := json.Unmarshal(accessTokenRespMarshel, &accessTokenResp); err != nil {
		// 	return err
		// }

		accessTokenResp := auth.AccessTokenResponse{
			AccessToken: qu["access_token"],
		}

		fmt.Println(accessTokenResp, "a, ", c.Queries())

		return c.JSON(fiber.Map{
			"access_token": accessTokenResp.AccessToken,
		})
	}
}

func SocialLoginRedirect() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Redirect("https://" + os.Getenv("AUTH0_DOMAIN") + "/authorize?response_type=token&client_id=" + os.Getenv("AUTH0_CLIENT_ID") + "&redirect_uri=http%3A%2F%2Flocalhost%3A4000%2Fauth%2Fcallback")
	}
}

func ChangePassword() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ChangePasswordReq := auth.ChangePasswordRequest{}
		if err := c.BodyParser(&ChangePasswordReq); err != nil {
			return err
		}

		passResponse, err := ChangePasswordReq.ChangePassword()
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"pass_response": passResponse,
		})
	}
}

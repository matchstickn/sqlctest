package auth

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type SignUpRequest struct {
	ClientId     string          `json:"client_id"`
	Email        string          `json:"email"`
	Password     string          `json:"password"`
	Connection   string          `json:"connection"`
	Username     string          `json:"username"`
	GivenName    string          `json:"given_name"`
	FamilyName   string          `json:"family_name"`
	Name         string          `json:"name"`
	Nickname     string          `json:"nickname"`
	Picture      string          `json:"picture"`
	UserMetadata json.RawMessage `json:"user_metadata"`
}

type SignUpResponse struct {
	GivenName     string          `json:"given_name"`
	FamilyName    string          `json:"family_name"`
	Name          string          `json:"name"`
	Nickname      string          `json:"nickname"`
	Picture       string          `json:"picture"`
	Id            string          `json:"_id"`
	Email         string          `json:"email"`
	Username      string          `json:"username"`
	EmailVerified bool            `json:"email_verified"`
	UserMetadata  json.RawMessage `json:"user_metadata"`
}

func (s SignUpRequest) GetSignUpId() (string, bool, error) {
	jsonBody, err := json.Marshal(s)
	if err != nil {
		return "", false, err
	}

	readerBody := bytes.NewReader(jsonBody)

	resp, err := http.Post("https://"+os.Getenv("AUTH0_DOMAIN")+"/dbconnections/signup", "application/json", readerBody)
	if err != nil {
		return "", false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", false, err
	}

	signUpData := SignUpResponse{}
	if err = json.Unmarshal(body, &signUpData); err != nil {
		return "", false, err
	}

	return signUpData.Id, signUpData.EmailVerified, nil
}

type AccessTokenRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Audience     string `json:"audience"`
	GrantType    string `json:"grant_type"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (a AccessTokenRequest) GetAccessToken() (string, error) {
	jsonBody, err := json.Marshal(a)
	if err != nil {
		return "", err
	}

	readerBody := bytes.NewReader(jsonBody)

	resp, err := http.Post("https://"+os.Getenv("AUTH0_DOMAIN")+"/oauth/token", "application/json", readerBody)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	AccessTokenResponse := AccessTokenResponse{}
	if err = json.Unmarshal(body, &AccessTokenResponse); err != nil {
		return "", err
	}

	return AccessTokenResponse.AccessToken, nil
}

package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"gitlab.com/moneropay/metronero/metronero-backend/app/models"

	"gitlab.com/moneropay/metronero/metronero-frontend/utils/config"
)

var (
	ErrUnauthorized = errors.New("Username or password is wrong.")
	ErrRegisterFail = errors.New("Failed to register new user.")
)

func UserLogin(username, password string) (*models.TokenInfo, error) {
	endpoint, err := url.JoinPath(config.Backend, "/login")
	if err != nil {
		return nil, err
	}

	resp, err := http.PostForm(endpoint, url.Values{
		"username": []string{username},
		"password": []string{password},
	})
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrUnauthorized
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var token models.TokenInfo
	err = json.Unmarshal(b, &token)
	return &token, err
}

func UserRegister(username, password string) error {
	endpoint, err := url.JoinPath(config.Backend, "/register")
	if err != nil {
		return err
	}

	resp, err := http.PostForm(endpoint, url.Values{
		"username": []string{username},
		"password": []string{password},
	})
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return ErrRegisterFail
	}
	return nil
}

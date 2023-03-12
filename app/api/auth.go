package api

import (
	"net/http"
	"net/url"
	"io"
	"encoding/json"

	"gitlab.com/moneropay/metronero/metronero-backend/app/models"

	"gitlab.com/moneropay/metronero/metronero-frontend/utils/config"
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

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var token models.TokenInfo
	err = json.Unmarshal(b, &token)
	return &token, err
}

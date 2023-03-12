package api

import (
	"net/http"
	"net/url"
	"fmt"
	"io"
	"encoding/json"

	"github.com/golang-jwt/jwt/v4"

	"gitlab.com/moneropay/metronero/metronero-backend/app/models"

	"gitlab.com/moneropay/metronero/metronero-frontend/utils/config"
)

func backendRequest(token *jwt.Token, method string) ([]byte, error) {
	endpoint, err := url.JoinPath(config.Backend, method)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token.Raw))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(res.Body)
}

func GetAdminDashboard(token *jwt.Token) (*models.AdminDashboardInfo, error) {
	resp, err := backendRequest(token, "/admin")
	if err != nil {
		return nil, err
	}
	var a models.AdminDashboardInfo
	err = json.Unmarshal(resp, &a)
	return &a, err
}

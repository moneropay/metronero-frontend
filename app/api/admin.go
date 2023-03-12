package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"gitlab.com/moneropay/metronero/metronero-backend/app/models"

	"gitlab.com/moneropay/metronero/metronero-frontend/utils/config"
)

func backendRequest(token, method string) ([]byte, error) {
	endpoint, err := url.JoinPath(config.Backend, method)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(res.Body)
}

func GetAdminDashboard(token string) (*models.AdminDashboardInfo, error) {
	resp, err := backendRequest(token, "/admin")
	if err != nil {
		return nil, err
	}
	var a models.AdminDashboardInfo
	err = json.Unmarshal(resp, &a)
	return &a, err
}

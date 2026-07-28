package tvdb

import (
	"anime-meta/lib/core/env"
	"bytes"
	"encoding/json"
	"fmt"

	"net/http"
)

func login() (string, error) {
	apiKey, err := env.GetVar("TVDB_APIKEY")
	if err != nil {
		return "", err
	}

	body, _ := json.Marshal(map[string]string{
		"apikey": apiKey,
	})

	resp, err := http.Post(BaseURL+"/login", "application/json", bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("tvdb login http call error: %w", err)
	}

	defer resp.Body.Close()

	var loginResp LoginResponse

	err = json.NewDecoder(resp.Body).Decode(&loginResp)
	if err != nil {
		return "", fmt.Errorf("tvdb login body parsing error: %w", err)
	}

	// if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
	// 	return "", fmt.Errorf("tvdb login body parsing error: %w", err)
	// }

	return loginResp.Data.Token, nil
}

func Authorize() (Client, error) {
	token, err := login()

	if err != nil {
		return Client{}, err
	}

	return Client{
		Token: token,
	}, nil
}

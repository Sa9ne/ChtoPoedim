package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type VerifyResponse struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Valid  bool   `json:"valid"`
}

func VerifyToken(token string) (*VerifyResponse, error) {
	url := "http://localhost:8081/verify"

	reqBody := []byte(fmt.Sprintf(`{"token":"%s"}`, token))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to verify token, status: %d", resp.StatusCode)
	}

	var response VerifyResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

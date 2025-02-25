package utils

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func VerifyRecaptcha(secretKey, recaptchaResponse string) (bool, error) {
	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify",
		url.Values{"secret": {secretKey}, "response": {recaptchaResponse}})
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var result struct {
		Success bool `json:"success"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}
	return result.Success, nil
}

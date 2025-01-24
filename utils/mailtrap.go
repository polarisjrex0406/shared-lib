package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type (
	EmailPayload struct {
		From struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"from"`
		To []struct {
			Email string `json:"email"`
		} `json:"to"`
		Subject  string `json:"subject"`
		Text     string `json:"text"`
		HTML     string `json:"html"`
		Category string `json:"category"`
	}
)

func SendEmail(mailTrapUrl, mailTrapApiKey string, payload EmailPayload) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", mailTrapUrl, strings.NewReader(string(jsonPayload)))

	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+mailTrapApiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		errorMsgs := struct {
			Errors []string `json:"errors"`
		}{}
		err = json.Unmarshal(body, &errorMsgs)
		if err != nil {
			return err
		}

		err = errors.New("")
		for _, msg := range errorMsgs.Errors {
			err = errors.Join(fmt.Errorf("%s", strings.ToLower(msg)))
		}
		return err
	}

	return nil
}

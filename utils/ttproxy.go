package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type TTProxySublicense struct {
	ID           uint   `json:"id"`
	Key          string `json:"key"`
	Secret       string `json:"secret"`
	ObtainLimit  int    `json:"obtainLimit"`
	TrafficLeft  int64  `json:"trafficLeft"`
	IPDuration   int    `json:"ipDuration"`
	Remark       string `json:"remark"`
	TotalTraffic int64  `json:"totalTraffic"`
	IPUsed       int    `json:"ipUsed"`
}

func TTProxyQueries(license, secret string, page int) url.Values {
	// Time to timestamp
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	// Add variables to query
	queries := url.Values{}
	queries.Set("license", license)
	queries.Set("time", ts)
	signBytes := md5.Sum([]byte(license + ts + secret))
	queries.Set("sign", hex.EncodeToString(signBytes[:]))
	// Add page if greater than 0
	if page > 0 {
		queries.Set("page", fmt.Sprintf("%d", page))
	}
	return queries
}

func TTProxyAPICall(method, baseUrl, endpoint string, queries url.Values, data url.Values) ([]byte, error) {
	payload := data.Encode()

	req, err := http.NewRequest(
		method,
		baseUrl+endpoint+"?"+queries.Encode(),
		strings.NewReader(payload),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body as a string
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API call failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	return bodyBytes, nil
}

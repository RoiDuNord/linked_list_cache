package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func InfoRequest() (string, string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/info?numbers=2,789,7,59,neflefjl,fhgfg,eong5,egv", nil)
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return "", ""
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Chrome/131.0.6778.86")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
		return "", ""
	}

	body := resp.Body
	status := resp.Status
	defer body.Close()

	var responseBody bytes.Buffer
	if _, err := io.Copy(&responseBody, body); err != nil {
		fmt.Println("Failed to read response body:", err)
		return "", ""
	}

	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		if err := json.Unmarshal(responseBody.Bytes(), &errorResponse); err != nil {
			fmt.Println("Failed to unmarshal error response:", err)
			return status, responseBody.String()
		}

		return status, fmt.Sprintf("{\"error\":\"%s\"}", errorResponse.Error)
	}

	return status, responseBody.String()
}

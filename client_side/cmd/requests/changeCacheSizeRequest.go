package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/server/handlers"
)

func ChangeCacheSizeRequest() (string, string) {
	client := http.Client{}
	url := "http://localhost:8080/admin/cache/changeSize"

	newSize := handlers.NewCacheSize{Size: 5}
	jsonData, err := json.Marshal(newSize)
	if err != nil {
		fmt.Println("Failed to marshal request body:", err)
		return "", ""
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return "", ""
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Chrome/131.0.6778.86")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
		return "", ""
	}
	defer resp.Body.Close()

	var responseBody bytes.Buffer
	if _, err := io.Copy(&responseBody, resp.Body); err != nil {
		fmt.Println("Failed to read response body:", err)
		return "", ""
	}

	status := resp.Status

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

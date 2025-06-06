package requests

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func sendRequest(method, url string, reqBody any) (string, any, int64) {
	start := time.Now()
	var reqData *bytes.Buffer

	if reqBody != nil {
		data, err := json.Marshal(reqBody)
		if err != nil {
			log.Fatalf("request body marshalling error: %v", err)
		}
		reqData = bytes.NewBuffer(data)
	} else {
		reqData = bytes.NewBuffer([]byte{})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, method, url, reqData)
	if err != nil {
		log.Fatalf("request error: %v", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Chrome/131.0.6778.86")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("response error %v", err)
	}

	body := resp.Body
	defer body.Close()

	var responseBody any
	if err := json.NewDecoder(body).Decode(&responseBody); err != nil {
		log.Fatalf("response body decoding error: %v", err)
	}

	responseTime := time.Since(start).Nanoseconds()

	return resp.Status, responseBody, responseTime
}

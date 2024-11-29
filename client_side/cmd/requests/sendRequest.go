package requests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func sendRequest(method, url string, reqBody interface{}) (string, interface{}, int64) {
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

	req, err := http.NewRequest(method, url, reqData)
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

	var responseBody interface{}
	if err := json.NewDecoder(body).Decode(&responseBody); err != nil {
		log.Fatalf("response body decoding error: %v", err)
	}

	responseTime := time.Since(start).Milliseconds()

	return resp.Status, responseBody, responseTime
}

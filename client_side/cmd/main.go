package main

import (
	"client/cmd/requests"
	"encoding/json"
	"fmt"
	"sync"
)

type Response struct {
	Endpoint string      `json:"endpoint"`
	Status   string      `json:"status"`
	Body     interface{} `json:"body"`
	Time     int64       `json:"responseTime(ms)"`
}

func executeRequest(name string, requestFunc func() (string, interface{}, int64)) Response {
	status, respBody, responseTime := requestFunc()
	return Response{
		Endpoint: name,
		Status:   status,
		Body:     respBody,
		Time:     responseTime,
	}
}

func main() {
	var wg sync.WaitGroup
	results := make([]Response, 0)

	funcs := []struct {
		name        string
		requestFunc func() (string, interface{}, int64)
	}{
		{"Info", requests.Info},
		{"ChangeCacheSize", requests.ChangeCacheSize},
		{"CacheOutputWithNewSize", requests.CacheOutput},
		{"ChangeFactor", requests.ChangeFactor},
		{"isActiveWorker", requests.IsActiveWorker},
		{"CacheOutputWithNewFactor", requests.CacheOutput},
		{"isActiveWorker", requests.IsActiveWorker},
	}

	for _, f := range funcs {
		wg.Add(1)
		go func(name string, requestFunc func() (string, interface{}, int64)) {
			defer wg.Done()
			result := executeRequest(name, requestFunc)
			results = append(results, result)
		}(f.name, f.requestFunc)
		wg.Wait()
	}

	for i, res := range results {
		jsonData, _ := json.MarshalIndent(res, "", "  ")
		if i != len(results)-1 {
			fmt.Println(string(jsonData) + "\n")
		} else {
			fmt.Print(string(jsonData) + "\n")
		}
	}
}

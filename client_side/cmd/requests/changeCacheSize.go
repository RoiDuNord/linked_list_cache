package requests

import (
	"server/server/handlers"
)

func ChangeCacheSize() (string, any, int64) {
	url := "http://localhost:8080/admin/cache/changeSize"
	newSize := handlers.NewCacheSize{Value: 2}

	status, respBody, responseTime := sendRequest("POST", url, newSize)

	return status, respBody, responseTime
}

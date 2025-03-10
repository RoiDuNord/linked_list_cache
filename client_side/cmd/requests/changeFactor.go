package requests

import (
	"server/server/handlers"
)

func ChangeFactor() (string, any, int64) {
	url := "http://localhost:8080/admin/cache/changeFactor"
	newFactor := handlers.Factor{Value: 20}

	status, respBody, responseTime := sendRequest("POST", url, newFactor)

	return status, respBody, responseTime
}

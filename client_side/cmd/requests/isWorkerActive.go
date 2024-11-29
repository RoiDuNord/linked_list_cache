package requests

func IsActiveWorker() (string, interface{}, int64) {
	url := "http://localhost:8080/admin/worker/setActiveStatus"

	status, respBody, time := sendRequest("POST", url, nil)

	return status, respBody, time
}

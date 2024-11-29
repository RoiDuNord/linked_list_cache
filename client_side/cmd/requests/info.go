package requests

func Info() (string, interface{}, int64) {
	url := "http://localhost:8080/info?numbers=2,789,7,59,neflefjl,fhgfg,eong5,egv"

	status, respBody, responseTime := sendRequest("GET", url, nil)

	return status, respBody, responseTime
}

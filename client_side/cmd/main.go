package main

import (
	"client/cmd/requests"
	"fmt"
)

func main() {
	status, response := requests.InfoRequest()
	fmt.Printf("%s\n%s", status, response)
}

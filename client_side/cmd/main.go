package main

import (
	"client/cmd/requests"
	"fmt"
)

func main() {
	status, response := requests.InfoRequest()
	fmt.Println("InfoRequest")
	fmt.Printf("%s\n%s", status, response)
	status1, response1 := requests.ChangeCacheSizeRequest()
	fmt.Println("\nChangeCacheSizeRequest")
	fmt.Printf("%s\n%s", status1, response1)
}

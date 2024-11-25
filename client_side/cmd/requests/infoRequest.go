package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"server/server/handlers"
	"strings"
)

func InfoRequest() {
	infoRequest := "GET /info?number=2,54,36,4,23,ejbcf,9,789,7,56,65,egv HTTP/1.1\n" +
		"Host: localhost:8080\n" +
		"Connection: close\n\n"

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	if _, err = conn.Write([]byte(infoRequest)); err != nil {
		fmt.Println("Error writing request:", err)
		return
	}

	response, err := io.ReadAll(conn)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	stringResponse := string(response)

	stringResponseSlice := strings.Split(stringResponse, "\r\n\r\n")

	var userResponseLine string
	for _, line := range stringResponseSlice {
		if strings.HasPrefix(line, "{") {
			userResponseLine = line
			break
		}
	}

	var userResponse handlers.UserResponse
	if err := json.Unmarshal([]byte(userResponseLine), &userResponse); err != nil {
		fmt.Println(err)
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	fmt.Println("UserResponse:", userResponse.Data)
}

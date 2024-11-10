package checkers

import (
	"fmt"
	"log"
	"net/http"
)

func Database(err error, w http.ResponseWriter) bool {
	if err != nil {
		log.Printf("%v", err)
		fmt.Fprintln(w, err)
		return false
	}
	return true
}

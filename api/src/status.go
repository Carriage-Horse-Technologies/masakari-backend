package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleGetServerStatus(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		status := 405
		w.WriteHeader(status)
		err := fmt.Errorf("This endpoint allows only GET method but recieve %s", r.Method)
		w.Write([]byte(err.Error()))

		LoggingHTTPError(status, err)
		return
	}

	result_json, err := getServerStatus()

	if err != nil {
		status := 500
		log.Println(err)
		w.WriteHeader(status)
		w.Write([]byte("Internal Server Error! Details in the log."))
		LoggingHTTPError(status, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result_json)
	LoggingHTTPError(200, fmt.Errorf("OK"))
}

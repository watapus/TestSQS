package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRouter() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/correspondence/fis", CreateMessageFromFIS).Methods("POST", "OPTIONS")

	return r
}

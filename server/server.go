package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodrigodealer/realtime/handlers"
)

func Server() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/subscription", handlers.SubscriptionHandler).Name("/subscription").Methods("GET")
	r.HandleFunc("/user", handlers.UserHandler).Name("/user").Methods("GET")
	return r
}

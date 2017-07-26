package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodrigodealer/realtime/es"
	"github.com/rodrigodealer/realtime/handlers"
)

func Server() http.Handler {
	client := es.Connect()
	es.IndexSetup(client, "facebook")

	r := mux.NewRouter()
	r.HandleFunc("/subscription", handlers.SubscriptionHandler).Name("/subscription").Methods("GET")
	r.HandleFunc("/subscription", handlers.FacebookUpdateHandler(client)).Name("/subscription").Methods("POST")
	r.HandleFunc("/user", handlers.UserHandler).Name("/user").Methods("GET")
	r.HandleFunc("/healthcheck", handlers.HealthcheckHandler).Name("/healthcheck").Methods("GET")
	return r
}

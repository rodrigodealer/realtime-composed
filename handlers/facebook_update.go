package handlers

import (
	"fmt"
	"net/http"

	"github.com/rodrigodealer/realtime/es"
	"github.com/rodrigodealer/realtime/models"
	"github.com/rodrigodealer/realtime/services"
)

func FacebookUpdateHandler(connection es.ElasticSearch,
	fbclient services.FacebookClient) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			var facebookUpdate = models.FacebookUpdate{}
			facebookUpdate.FromRequest(r.Body)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello, %v", r.URL.Query().Get("u"))
	}
}

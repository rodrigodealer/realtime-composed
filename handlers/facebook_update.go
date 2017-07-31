package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rodrigodealer/realtime/es"
	"github.com/rodrigodealer/realtime/models"
)

func FacebookUpdateHandler(connection es.ElasticSearch) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			var facebookUpdate = models.FacebookUpdate{}
			facebookUpdate.FromRequest(r.Body)
			log.Print(facebookUpdate)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello, %v", r.URL.Query().Get("u"))
	}
}

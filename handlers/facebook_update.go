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
			token := services.GetToken(fbclient)
			for _, entry := range facebookUpdate.Entry {
				var response = services.GetUid(entry, token, fbclient)
				facebookUser := &models.FacebookUser{}
				facebookUser.FromJson(response.Body)

				go connection.IndexUser("facebook", facebookUser)
			}
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello, %v", r.URL.Query().Get("u"))
	}
}

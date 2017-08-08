package handlers

import (
	"fmt"
	"net/http"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/rodrigodealer/realtime/es"
	"github.com/rodrigodealer/realtime/models"
	"github.com/rodrigodealer/realtime/services"
)

func FacebookUpdateHandler(connection es.ElasticSearch,
	fbclient services.FacebookClient) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			span := opentracing.StartSpan("Facebook API")
			var facebookUpdate = models.FacebookUpdate{}
			facebookUpdate.FromRequest(r.Body)
			token := services.GetToken(fbclient, span)
			for _, entry := range facebookUpdate.Entry {
				var response = services.GetUid(entry, token, fbclient, span)
				facebookUser := &models.FacebookUser{}
				facebookUser.FromJson(response.Body)

				go connection.IndexUser("facebook", facebookUser, span)
			}
			defer span.Finish()
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello, %v", r.URL.Query().Get("u"))
	}
}

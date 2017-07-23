package handlers

import (
	"fmt"
	"net/http"
)

func FacebookUpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %v", r.URL.Query().Get("u"))
}

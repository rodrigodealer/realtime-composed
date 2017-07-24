package handlers

import (
	"fmt"
	"net/http"
)

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "WORKING")
}

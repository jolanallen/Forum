package handler

import (
	"net/http"
	"fmt"
)

func OverloadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTooManyRequests)
	fmt.Fprintf(w, "<h1>503 Service Unavailable</h1><p>trop de reqest merci de patienter .</p>")
}

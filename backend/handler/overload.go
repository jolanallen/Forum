package handler

import (
	"net/http"
)

func OverloadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTooManyRequests)
	w.Write([]byte("Too many requests, please try again later."))
}

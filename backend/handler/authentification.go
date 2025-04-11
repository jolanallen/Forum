package handler

import (
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

//////////////dans validationService.go/////////////////////

}

func Register(w http.ResponseWriter, r *http.Request) {

//////////////dans registerService.go/////////////////////

	fmt.Fprintln(w, "Page Register")
}

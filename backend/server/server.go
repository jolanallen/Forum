package server

import (
	"fmt"
	"log"
	"net/http"
	"Forum/backend/utils"
)

func Server() {
	certFile := "backend/server/ssl_tls/cert.pem"
	keyFile := "backend/server/ssl_tls/key.pem"
	InitRoutes()

	

	fmt.Println("https://localhost:443")
	err := http.ListenAndServeTLS(":443", certFile, keyFile, utils.F.MainRouter)
	if err != nil {
		log.Fatal(err)
	}
}

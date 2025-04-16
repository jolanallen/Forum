package structs

import (
	"net/http"
)

// Forum represents the forum structure, containing the main router for handling HTTP requests.
// It is responsible for initializing and routing the different endpoints for the forum application.
type Forum struct {
	// MainRouter is the main HTTP router for the forum. It handles routing of HTTP requests to appropriate handlers.
	MainRouter *http.ServeMux
}


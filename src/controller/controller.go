package controller

import (
	"github.com/gorilla/mux"
)

var (
	githubController ghub
)

// Startup func to register routes
func Startup(r *mux.Router) {
	githubController.registerRoutes(r)
}

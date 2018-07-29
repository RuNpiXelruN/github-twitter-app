package controller

import (
	"go_apps/go_api_apps/github-twitter-app/src/github"
	"go_apps/go_api_apps/github-twitter-app/src/twitter"

	"github.com/gorilla/mux"
)

var (
	githubController github.Github

	// below for example only
	twitterController twitter.Twitter
)

// Startup func to register routes
func Startup(r *mux.Router) {
	githubController.RegisterRoutes(r)

	// twitterController.RegisterRoutes(r)  -- for example only
}

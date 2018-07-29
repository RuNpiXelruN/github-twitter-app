package main

import (
	"fmt"
	"go_apps/go_api_apps/github-twitter-app/config"
	"go_apps/go_api_apps/github-twitter-app/src/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	config.SetVars()
}

func main() {
	r := mux.NewRouter()
	controller.Startup(r)

	r.Path("/").HandlerFunc(index).Methods("GET")

	go http.ListenAndServe(":8080", r)
	go fmt.Println("..listening on port 8080")
	fmt.Scanln()
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Welcome to the Github/Twitter/Slack Go App")
}

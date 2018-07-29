package controller

import (
	"fmt"

	"github.com/gorilla/mux"
)

type tweet struct{}

func (t tweet) registerRoutes(r *mux.Router) {

}

// TweetIssue func
func TweetIssue(issTweet *IssueTweet) {
	fmt.Printf("\n%+v\n", issTweet)
}

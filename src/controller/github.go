package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
)

type ghub struct{}

// IssueTweet struct
type IssueTweet struct {
	ID             *int64  `json:"id"`
	Title          *string `json:"title"`
	URL            *string `json:"url"`
	Action         *string `json:"action"`
	User           *string `json:"user"`
	AvatarURL      *string `json:"avatar_url"`
	RepositoryName *string `json:"repository_name"`
}

func (g ghub) registerRoutes(r *mux.Router) {
	r.Path("/github-listener").HandlerFunc(g.githubEventHandler).Methods("POST")
}

func (g ghub) githubEventHandler(w http.ResponseWriter, req *http.Request) {
	payload, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error reading req Body:", err)
		return
	}
	defer req.Body.Close()

	event, err := github.ParseWebHook(github.WebHookType(req), payload)
	if err != nil {
		fmt.Println("Error parsing github webhook:", err)
	}

	// check the type of event coming from Github
	switch e := event.(type) {
	case *github.PullRequestEvent:
		fmt.Println("Pull request event recieved")
	case *github.IssuesEvent:
		handleGithubIssueEvent(e)
	default:
		fmt.Println("Something else happened")
	}
}

func handleGithubIssueEvent(e *github.IssuesEvent) {
	IssueDetails := e.Issue
	UserDetails := IssueDetails.User
	Repo := e.Repo
	Sender := e.Sender

	issue := IssueTweet{
		ID:             IssueDetails.ID,
		Title:          IssueDetails.Title,
		URL:            IssueDetails.URL,
		Action:         e.Action,
		User:           Sender.Login,
		AvatarURL:      UserDetails.AvatarURL,
		RepositoryName: Repo.Name,
	}

	TweetIssue(&issue)
}

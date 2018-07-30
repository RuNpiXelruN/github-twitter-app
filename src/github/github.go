package github

import (
	"fmt"
	"go_apps/go_api_apps/github-twitter-app/src/slack"
	"go_apps/go_api_apps/github-twitter-app/src/twitter"
	"io/ioutil"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
)

// Github struct
type Github struct{}

// RegisterRoutes func - register all github related endpoints
func (g Github) RegisterRoutes(r *mux.Router) {
	r.Path("/github-listener").HandlerFunc(g.githubEventHandler).Methods("POST")
}

// handle incoming github event and send to relevant func
func (g Github) githubEventHandler(w http.ResponseWriter, req *http.Request) {
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
		handleGithubPREvent(e)
	case *github.IssuesEvent:
		handleGithubIssueEvent(e)
	default:
		fmt.Println("An event we're not interested in occurred")
	}
}

func handleGithubIssueEvent(e *github.IssuesEvent) {
	if *e.Action == "opened" || *e.Action == "reopened" || *e.Action == "closed" {
		IssueDetails := e.Issue
		UserDetails := IssueDetails.User
		Repo := e.Repo
		Sender := e.Sender

		issue := twitter.IssueTweet{
			ID:             IssueDetails.ID,
			Title:          IssueDetails.Title,
			URL:            IssueDetails.URL,
			Action:         e.Action,
			User:           Sender.Login,
			AvatarURL:      UserDetails.AvatarURL,
			RepositoryName: Repo.Name,
		}

		twitter.SendIssueDetailsToTwitter(&issue)
	}
}

func handleGithubPREvent(e *github.PullRequestEvent) {
	if *e.Action == "opened" || *e.Action == "reopened" {
		PRDetails := e.PullRequest
		UserDetails := PRDetails.User
		Repo := e.Repo
		Sender := e.Sender

		pullRequest := slack.PRSlack{
			ID:             PRDetails.ID,
			Action:         e.Action,
			URL:            PRDetails.HTMLURL,
			Title:          PRDetails.Title,
			PRNumber:       PRDetails.Number,
			AvatarURL:      UserDetails.AvatarURL,
			User:           Sender.Login,
			RepositoryName: Repo.Name,
		}
		slack.SendPR(&pullRequest)
	}
}

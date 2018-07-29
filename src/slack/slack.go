package slack

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

var (
	api *slack.Client
)

// PRSlack type
type PRSlack struct {
	ID             *int64  `json:"id"`
	Action         *string `json:"action"`
	URL            *string `json:"url"`
	Title          *string `json:"title"`
	AvatarURL      *string `json:"avatar_url"`
	User           *string `json:"user"`
	RepositoryName *string `json:"repository_name"`
}

// SlackListener struct, overkill for a bot, but for a slack app
// would store reference to botID and channelID
type SlackListener struct {
	client *slack.Client
}

var (
	client        *slack.Client
	slackListener *SlackListener
)

// Startup func
func Startup() {
	client = slack.New(os.Getenv("SLACK-API-TOKEN"))

	slackListener = &SlackListener{
		client: client,
	}
}

// SendPR func handles posting of PR details to slack channel
func SendPR(pr *PRSlack) {
	groups, err := slackListener.client.GetGroups(false)
	if err != nil {
		fmt.Println("Error fetching groups->", err)
	}

	attachment := slack.Attachment{
		Text:       "_" + *pr.User + "_ has *" + *pr.Action + "* a pull request for *" + *pr.RepositoryName + "*\n" + *pr.URL,
		Color:      "#2196f3",
		Title:      "New PR",
		TitleLink:  *pr.URL,
		ThumbURL:   *pr.AvatarURL,
		Footer:     "created by me :)",
		FooterIcon: "https://user-images.githubusercontent.com/13185159/43365880-faded25a-9376-11e8-87dc-36aea54ac547.png",
	}

	params := slack.PostMessageParameters{
		Attachments: []slack.Attachment{
			attachment,
		},
	}
	params.AsUser = true

	for _, group := range groups {
		slackListener.client.PostMessage(group.Name, "A PR is up for review!!", params)
	}
}

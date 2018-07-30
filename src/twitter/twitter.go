package twitter

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Twitter struct{}

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

var mediaUploadPath = "https://upload.twitter.com/1.1/media/upload.json"

// register all twitter related endpoints
func (t Twitter) registerRoutes(r *mux.Router) {

}

// SendIssueDetailsToTwitter func
func SendIssueDetailsToTwitter(issTweet *IssueTweet) {

	resp, err := http.Get(*issTweet.AvatarURL)
	if err != nil {
		fmt.Println("Error fetching github avatar URL")
		return
	}
	defer resp.Body.Close()

	file, err := os.Create("twitterImg.jpg")
	if err != nil {
		fmt.Println("Error creating image file")
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error writing bytes to image file")
	}

	form := createInitMediaForm(file)
	req, err := http.NewRequest("POST", mediaUploadPath, strings.NewReader(*form))
	if err != nil {
		fmt.Println("Error sending request to twitter media url")
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// If a Twitter dev account was available i would
	// now send my request and expect to recieve a response payload
	// which included my `MediaID` needed for uploading images to twitter.

	// below for example cleanup only
	time.Sleep(3 * time.Second)
	os.Remove(file.Name())
}

// build up url values to init twitter media
func createInitMediaForm(f *os.File) *string {
	fileInfo, err := f.Stat()
	if err != nil {
		fmt.Println("Error extracting file info")
		return nil
	}

	form := url.Values{}
	form.Add("command", "INIT")
	form.Add("media_type", mime.TypeByExtension(filepath.Ext(f.Name())))
	form.Add("total_bytes", fmt.Sprint(fileInfo.Size()))

	initMediaString := form.Encode()
	return &initMediaString
}

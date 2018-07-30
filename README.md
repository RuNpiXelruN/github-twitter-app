# Github / Twitter / Slack Bot App
This App is an event-driven http service built in Go.
This App revieves data from a given Github repository via a webhook and currently is listening for New & Close Issues, as well as New & Reopened Pull Requests.  

## Slack
The Slack side of the App recieves New and Reopened Pull Request events and posts a link and details of the PR to a given slack channel.

## Twitter
The Twitter side of the App recieves New and Closed Issue events and posts and aims to post the Issue details to twitter however, at the time of writing Twitter has increased scrutiny around new apps and so my developer application is still pending. 

## Setup
### Github Setup
• To connect a Github repository to this app you need to either own or be an admin of the repository.  
• Go to the Github repository, click on settings, choose webhooks, and copy the url or this App into the webhook input field. Please ensure that both _Issues_ and _Pull Requests_ event types are selected. You can do this by selecting the _individual events_ radio button and then choosing these options, or selecting the _Send me everything_ radio button.  
• Once selected, save/update webhook at the bottom of the page. The Github element has now been setup.  
 ** If you want to test this app locally I recommend using a serve such as _ngrok_ to create a secure tunnel to your localhost, and then reference the _ngrok_ url as your github webhook url.
### Slack Setup
• To connect Slack to this App you'll need to find your _SLACK-API-TOKEN_ for your team and update the _SLACK-API-TOKEN_ config variable within the config folder of this app.
• You then need to simply add this Bot to any channel you wish to recieve Pull Request updates on.

## Get in touch
For any questions or troubleshooting please feel free to get in touch at justindavidson23@gmail.com.  
Pull Requests welcome :)
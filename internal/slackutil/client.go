package slackutil

import (
	"log"
	"os"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

var apiClient *slack.Client = nil
var botClient *socketmode.Client = nil

func Client(appToken string, botToken string) (*slack.Client, *socketmode.Client) {
	if apiClient == nil && appToken != "" {
		apiClient = slack.New(
			botToken,
			slack.OptionDebug(true),
			slack.OptionLog(log.New(os.Stdout, "api: ", log.Lshortfile|log.LstdFlags)),
			slack.OptionAppLevelToken(appToken),
		)
	}

	if botClient == nil && botToken != "" {
		botClient = socketmode.New(
			apiClient,
			socketmode.OptionDebug(true),
			socketmode.OptionLog(log.New(os.Stdout, "socket: ", log.Lshortfile|log.LstdFlags)),
		)
	}

	return apiClient, botClient
}

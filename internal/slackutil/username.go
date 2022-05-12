package slackutil

import (
	"github.com/slack-go/slack/socketmode"
)

func GetUsername(client *socketmode.Client, id string) (string, error) {
	user, err := client.GetUserInfo(id)
	if user != nil {
		return user.Name, err
	}
	return id, err
}

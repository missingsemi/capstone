package util

import (
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func UpdateView(client *socketmode.Client, event socketmode.Event, view slack.ModalViewRequest) {
	client.Ack(*event.Request, struct {
		ResponseAction string      `json:"response_action"`
		View           interface{} `json:"view"`
	}{
		ResponseAction: "update",
		View:           view,
	})
}

func UpdateView2(client *socketmode.Client, callback slack.InteractionCallback, view slack.ModalViewRequest) {
	client.UpdateView(view, "", "", callback.View.ID)
}

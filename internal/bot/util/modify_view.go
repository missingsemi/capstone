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

func UpdateView2(client *socketmode.Client, viewId string, view slack.ModalViewRequest) {
	client.UpdateView(view, "", "", viewId)
}

func PushView(client *socketmode.Client, event socketmode.Event, view slack.ModalViewRequest) {
	client.Ack(*event.Request, struct {
		ResponseAction string      `json:"response_action"`
		View           interface{} `json:"view"`
	}{
		ResponseAction: "push",
		View:           view,
	})
}

func PushView2(client *socketmode.Client, triggerId string, view slack.ModalViewRequest) {
	client.PushView(triggerId, view)
}

func ErrorView(client *socketmode.Client, event socketmode.Event, errors map[string]string) {
	client.Ack(*event.Request, struct {
		ResponseAction string            `json:"response_action"`
		Errors         map[string]string `json:"errors"`
	}{
		ResponseAction: "errors",
		Errors:         errors,
	})
}

package controller

import (
	"errors"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

type CallbackHandler = func(*socketmode.Client, slack.InteractionCallback) error

var callbackHandlers map[string]CallbackHandler

func RegisterCallbackHandler(callbackId string, cbh CallbackHandler) {
	callbackHandlers[callbackId] = cbh
}

func UnregisterCallbackHandler(callbackId string) {
	delete(callbackHandlers, callbackId)
}

func CallCallbackHandler(callbackId string, client *socketmode.Client, callback slack.InteractionCallback) error {
	if cbh, ok := callbackHandlers[callbackId]; ok {
		return cbh(client, callback)
	} else {
		return errors.New("No handler registered for callbackId \"" + callbackId + "\".")
	}
}

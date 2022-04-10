package controller

import (
	"errors"

	"github.com/slack-go/slack/socketmode"
)

type CallbackHandlerFunc = func(*socketmode.Client, socketmode.Event) error

var callbackHandlers map[string]CallbackHandlerFunc = make(map[string]CallbackHandlerFunc)

func RegisterCallbackHandler(callbackId string, fn CallbackHandlerFunc) {
	callbackHandlers[callbackId] = fn
}

func UnregisterCallbackHandler(callbackId string) {
	delete(callbackHandlers, callbackId)
}

func CallCallbackHandler(callbackId string, client *socketmode.Client, event socketmode.Event) error {
	if fn, ok := callbackHandlers[callbackId]; ok {
		return fn(client, event)
	} else {
		return errors.New("No handler registered for callbackId \"" + callbackId + "\".")
	}
}

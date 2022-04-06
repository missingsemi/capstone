package controller

import (
	"errors"

	"github.com/slack-go/slack/socketmode"
)

type CallbackHandlerFunc = func(*socketmode.Client, socketmode.Event, interface{}) error

type CallbackHandler struct {
	fn   CallbackHandlerFunc
	data interface{}
}

var callbackHandlers map[string]CallbackHandler = make(map[string]CallbackHandler)

func RegisterCallbackHandler(callbackId string, fn CallbackHandlerFunc, data interface{}) {
	callbackHandlers[callbackId] = CallbackHandler{
		fn,
		data,
	}
}

func UnregisterCallbackHandler(callbackId string) {
	delete(callbackHandlers, callbackId)
}

func CallCallbackHandler(callbackId string, client *socketmode.Client, event socketmode.Event) error {
	if cbh, ok := callbackHandlers[callbackId]; ok {
		return cbh.fn(client, event, cbh.data)
	} else {
		return errors.New("No handler registered for callbackId \"" + callbackId + "\".")
	}
}

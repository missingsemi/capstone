package controller

import (
	"errors"

	"github.com/slack-go/slack/socketmode"
)

type CommandHandlerFunc = func(*socketmode.Client, socketmode.Event, interface{}) error

type CommandHandler struct {
	fn   CommandHandlerFunc
	data interface{}
}

var commandHandlers map[string]CommandHandler = make(map[string]CommandHandler)

func RegisterCommandHandler(commandId string, fn CommandHandlerFunc, data interface{}) {
	commandHandlers[commandId] = CommandHandler{
		fn,
		data,
	}
}

func UnregisterCommandHandler(commandId string) {
	delete(commandHandlers, commandId)
}

func CallCommandHandler(commandId string, client *socketmode.Client, event socketmode.Event) error {
	if cbh, ok := commandHandlers[commandId]; ok {
		return cbh.fn(client, event, cbh.data)
	} else {
		return errors.New("No handler registered for commandId \"" + commandId + "\".")
	}
}

package controller

import (
	"errors"

	"github.com/slack-go/slack/socketmode"
)

type CommandHandlerFunc = func(*socketmode.Client, socketmode.Event) error

var commandHandlers map[string]CommandHandlerFunc = make(map[string]CommandHandlerFunc)

func RegisterCommandHandler(commandId string, fn CommandHandlerFunc) {
	commandHandlers[commandId] = fn
}

func UnregisterCommandHandler(commandId string) {
	delete(commandHandlers, commandId)
}

func CallCommandHandler(commandId string, client *socketmode.Client, event socketmode.Event) error {
	if fn, ok := commandHandlers[commandId]; ok {
		return fn(client, event)
	} else {
		return errors.New("No handler registered for commandId \"" + commandId + "\".")
	}
}

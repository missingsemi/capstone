package controller

import (
	"errors"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

type CommandHandler = func(*socketmode.Client, slack.SlashCommand) error

var commandHandlers map[string]CommandHandler

func RegisterCommandHandler(command string, cmdh CommandHandler) {
	commandHandlers[command] = cmdh
}

func UnregisterCommandHandler(command string) {
	delete(commandHandlers, command)
}

func CallCommandHandler(command string, client *socketmode.Client, slashCommand slack.SlashCommand) error {
	if cmdh, ok := commandHandlers[command]; ok {
		return cmdh(client, slashCommand)
	} else {
		return errors.New("No handler registered for command \"" + command + "\".")
	}
}

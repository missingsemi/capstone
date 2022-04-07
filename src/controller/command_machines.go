package controller

import (
	"errors"

	"github.com/missingsemi/capstone/database"
	"github.com/missingsemi/capstone/view"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func CommandMachines(client *socketmode.Client, event socketmode.Event, data interface{}) error {

	command, ok := event.Data.(slack.SlashCommand)
	if !ok {
		return errors.New("type assertion failed")
	}

	machines, err := database.GetMachines()
	if err != nil {
		return err
	}

	modalView := view.AdminMachinesAvailableMachines(machines)

	_, err = client.OpenView(command.TriggerID, modalView)
	if err != nil {
		return err
	}

	client.Ack(*event.Request)
	return nil
}
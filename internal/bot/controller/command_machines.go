package controller

import (
	"errors"

	"github.com/missingsemi/capstone/internal/bot/view"
	"github.com/missingsemi/capstone/internal/database"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func CommandMachines(client *socketmode.Client, event socketmode.Event) error {

	command, ok := event.Data.(slack.SlashCommand)
	if !ok {
		return errors.New("type assertion failed")
	}

	machines, err := database.GetMachines()
	if err != nil {
		return err
	}

	isAdmin, _ := database.IsUserAdmin(command.UserID)

	var modalView slack.ModalViewRequest

	if isAdmin {
		modalView = view.AdminMachinesAvailableMachines(machines)
	} else {
		modalView = view.UserMachinesAvailableMachines(machines)
	}

	_, err = client.OpenView(command.TriggerID, modalView)
	if err != nil {
		return err
	}

	client.Ack(*event.Request)
	return nil
}

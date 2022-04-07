package controller

import (
	"errors"

	"github.com/missingsemi/capstone/database"
	"github.com/missingsemi/capstone/util"
	"github.com/missingsemi/capstone/view"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func CallbackAvailableMachines(client *socketmode.Client, event socketmode.Event, data interface{}) error {
	callback, ok := event.Data.(slack.InteractionCallback)
	if !ok {
		return errors.New("type assertion failed")
	}

	actions := callback.ActionCallback.BlockActions
	if len(actions) == 0 {
		client.Ack(*event.Request)
		return errors.New("no action provided")
	} else if actions[0].Type != "button" {
		client.Ack(*event.Request)
		return errors.New("action was of wrong type")
	}

	if actions[0].ActionID == "admin_machines-available_machines-delete_callback" {
		err := database.DeleteMachine(actions[0].Value)
		if err != nil {
			client.Ack(*event.Request)
			return err
		}

		machines, err := database.GetMachines()

		util.UpdateView2(client, callback, view.AdminMachinesAvailableMachines(machines))
		client.Ack(*event.Request)

		return err
	}

	client.Ack(*event.Request)
	return nil
}

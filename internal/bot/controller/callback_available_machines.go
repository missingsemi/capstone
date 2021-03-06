package controller

import (
	"errors"

	"github.com/missingsemi/capstone/internal/bot/view"
	"github.com/missingsemi/capstone/internal/database"
	"github.com/missingsemi/capstone/internal/slackutil"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func CallbackAvailableMachines(client *socketmode.Client, event socketmode.Event) error {
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

		slackutil.UpdateView2(client, callback.View.ID, view.AdminMachinesAvailableMachines(machines))
		client.Ack(*event.Request)

		return err
	}

	if actions[0].ActionID == "admin_machines-available_machines-edit_callback" {
		machine, err := database.GetMachineById(actions[0].Value)
		if err != nil {
			client.Ack(*event.Request)
			return err
		}

		slackutil.PushView2(client, callback.TriggerID, view.AdminMachinesEditMachine(machine))
		client.Ack(*event.Request)

		return err
	}

	if actions[0].ActionID == "admin_machines-available_machines-create_callback" {
		slackutil.PushView2(client, callback.TriggerID, view.AdminMachinesCreateMachine())
		client.Ack(*event.Request)
	}

	client.Ack(*event.Request)
	return nil
}

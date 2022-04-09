package controller

import (
	"errors"

	"github.com/missingsemi/capstone/database"
	"github.com/missingsemi/capstone/view"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func CommandSchedule(client *socketmode.Client, event socketmode.Event, data interface{}) error {
	command, ok := event.Data.(slack.SlashCommand)
	if !ok {
		return errors.New("type assertion failed")
	}

	sessions, err := database.GetUpcomingSessionsByUser(command.UserID)
	if err != nil {
		return err
	}

	machines, err := database.GetMachines()
	if err != nil {
		return err
	}

	modalView := view.UserScheduleCreatedSessions(sessions, machines)

	_, err = client.OpenView(command.TriggerID, modalView)
	if err != nil {
		return err
	}

	client.Ack(*event.Request)
	return nil
}

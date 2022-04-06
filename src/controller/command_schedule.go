package controller

import (
	"errors"

	"github.com/missingsemi/capstone/database"
	"github.com/missingsemi/capstone/model"
	"github.com/missingsemi/capstone/view"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func CommandSchedule(client *socketmode.Client, event socketmode.Event, data interface{}) error {
	sessions, ok := data.(map[string]model.ScheduleSession)
	if !ok {
		return errors.New("type assertion failed")
	}

	command, ok := event.Data.(slack.SlashCommand)
	if !ok {
		return errors.New("type assertion failed")
	}

	machines, err := database.GetMachines()
	if err != nil {
		return err
	}

	modalView := view.ScheduleAddTeamInformation(machines)
	_, err = client.OpenView(command.TriggerID, modalView)
	if err != nil {
		return err
	}

	session := model.ScheduleSession{}
	session.UserId = command.UserID
	sessions[command.UserID] = session

	client.Ack(*event.Request)
	return nil
}

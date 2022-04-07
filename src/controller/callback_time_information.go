package controller

import (
	"errors"
	"time"

	"github.com/missingsemi/capstone/database"
	"github.com/missingsemi/capstone/model"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func CallbackTimeInformation(client *socketmode.Client, event socketmode.Event, data interface{}) error {
	sessions, ok := data.(map[string]model.ScheduleSession)
	if !ok {
		return errors.New("type assertion failed")
	}

	callback, ok := event.Data.(slack.InteractionCallback)
	if !ok {
		return errors.New("type assertion failed")
	}

	timeStr := callback.View.State.Values["time_input_block"]["time_input"].SelectedOption.Value
	session := sessions[callback.User.ID]

	session.Time, _ = time.Parse(time.RFC3339, timeStr)

	database.CreateSession(session)
	delete(sessions, callback.User.ID)

	client.Ack(*event.Request)

	return nil
}

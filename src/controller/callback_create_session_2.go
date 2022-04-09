package controller

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/missingsemi/capstone/database"
	"github.com/missingsemi/capstone/model"
	"github.com/missingsemi/capstone/util"
	"github.com/missingsemi/capstone/view"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func CallbackCreateSession2(client *socketmode.Client, event socketmode.Event, data interface{}) error {
	callback, ok := event.Data.(slack.InteractionCallback)
	if !ok {
		return errors.New("type assertion failed")
	}

	timeStr := callback.View.State.Values["time_input_block"]["time_input"].SelectedOption.Value

	var session model.ScheduleSession
	json.Unmarshal([]byte(callback.View.PrivateMetadata), &session)

	session.Time, _ = time.Parse(time.RFC3339, timeStr)

	database.CreateSession(session)

	sessions, _ := database.GetUpcomingSessionsByUser(session.UserId)
	machines, _ := database.GetMachines()
	util.UpdateView2(client, callback.View.PreviousViewID, view.UserScheduleCreatedSessions(sessions, machines))
	client.Ack(*event.Request)
	return nil
}
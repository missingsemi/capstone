package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/missingsemi/capstone/internal/bot/view"
	"github.com/missingsemi/capstone/internal/database"
	"github.com/missingsemi/capstone/internal/slackutil"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func CallbackCreatedSessions(client *socketmode.Client, event socketmode.Event) error {
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

	if actions[0].ActionID == "user_schedule-created_sessions-delete_callback" {
		sessionId, _ := strconv.ParseInt(actions[0].Value, 10, 32)

		err := database.DeleteSession(int(sessionId))
		if err != nil {
			client.Ack(*event.Request)
			return err
		}

		sessions, err := database.GetUpcomingSessionsByUser(callback.User.ID)
		machines, err := database.GetMachines()

		slackutil.UpdateView2(client, callback.View.ID, view.UserScheduleCreatedSessions(sessions, machines))
		client.Ack(*event.Request)

		return err
	}
	if actions[0].ActionID == "user_schedule-created_sessions-view_callback" {
		sessionId, _ := strconv.ParseInt(actions[0].Value, 10, 32)

		session, err := database.GetSessionById(int(sessionId))
		if err != nil {
			client.Ack(*event.Request)
			return err
		}
		machine, _ := database.GetMachineById(session.Machine)

		slackutil.PushView2(client, callback.TriggerID, view.UserMachinesViewSession(session, machine))
		client.Ack(*event.Request)

		return err
	}

	if actions[0].ActionID == "user_schedule-created_sessions-create_callback" {
		machines, _ := database.GetMachines()
		v := view.UserScheduleCreateSession1(machines)
		j, _ := json.Marshal(v)
		fmt.Println(string(j))
		slackutil.PushView2(client, callback.TriggerID, view.UserScheduleCreateSession1(machines))
		client.Ack(*event.Request)
	}

	client.Ack(*event.Request)
	return nil
}

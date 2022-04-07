package controller

import (
	"errors"
	"strconv"

	"github.com/missingsemi/capstone/database"
	"github.com/missingsemi/capstone/model"
	"github.com/missingsemi/capstone/util"
	"github.com/missingsemi/capstone/view"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func CallbackMachineInformation(client *socketmode.Client, event socketmode.Event, data interface{}) error {
	sessions, ok := data.(map[string]model.ScheduleSession)
	if !ok {
		return errors.New("type assertion failed")
	}

	callback, ok := event.Data.(slack.InteractionCallback)
	if !ok {
		return errors.New("type assertion failed")
	}

	duration, _ := strconv.ParseInt(callback.View.State.Values["duration_input_block"]["duration_input"].SelectedOption.Value, 10, 32)
	session := sessions[callback.User.ID]
	session.Reason = callback.View.State.Values["reason_input_block"]["reason_input"].Value
	session.Duration = duration
	sessions[callback.User.ID] = session

	machine, _ := database.GetMachineById(session.Machine)
	// Only have to consider sessions that could still be ongoing at the current time, so time.Now() - 4h
	validTimes := util.GenerateValidTimes(machine, int(duration))

	util.UpdateView(client, event, view.ScheduleAddTimeInformation(validTimes))

	return nil
}

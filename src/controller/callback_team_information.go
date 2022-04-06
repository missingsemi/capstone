package controller

import (
	"errors"

	"github.com/missingsemi/capstone/database"
	"github.com/missingsemi/capstone/model"
	"github.com/missingsemi/capstone/util"
	"github.com/missingsemi/capstone/view"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func CallbackTeamInformation(client *socketmode.Client, event socketmode.Event, data interface{}) error {
	sessions, ok := data.(map[string]model.ScheduleSession)
	if !ok {
		return errors.New("type assertion failed")
	}

	callback, ok := event.Data.(slack.InteractionCallback)
	if !ok {
		return errors.New("type assertion failed")
	}

	machineId := callback.View.State.Values["machine_input_block"]["machine_input"].SelectedOption.Value

	session := sessions[callback.User.ID]
	session.GroupIds = callback.View.State.Values["teammates_input_block"]["teammates_input"].SelectedUsers
	session.Machine = machineId
	sessions[callback.User.ID] = session

	machine, _ := database.GetMachineById(machineId)
	util.UpdateView(client, event, view.ScheduleAddMachineInformation(machine))
	return nil
}

package controller

import (
	"errors"
	"strconv"

	"github.com/missingsemi/capstone/internal/bot/util"
	"github.com/missingsemi/capstone/internal/bot/view"
	"github.com/missingsemi/capstone/pkg/database"
	"github.com/missingsemi/capstone/pkg/model"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func CallbackCreateSession1(client *socketmode.Client, event socketmode.Event) error {
	callback, ok := event.Data.(slack.InteractionCallback)
	if !ok {
		return errors.New("type assertion failed")
	}

	groupInput := callback.View.State.Values["group_input_block"]["group_input"].SelectedUsers
	machineInput := callback.View.State.Values["machine_input_block"]["machine_input"].SelectedOption.Value
	reasonInputBlock := callback.View.State.Values["reason_input_block"]["reason_input"].Value
	durationInputBlock, _ := strconv.ParseInt(callback.View.State.Values["duration_input_block"]["duration_input"].SelectedOption.Value, 10, 32)

	session := model.ScheduleSession{
		UserId:   callback.User.ID,
		GroupIds: groupInput,
		Machine:  machineInput,
		Reason:   reasonInputBlock,
		Duration: durationInputBlock,
	}

	machine, _ := database.GetMachineById(session.Machine)

	validTimes := util.GenerateValidTimes(machine, int(durationInputBlock))

	util.UpdateView(client, event, view.UserScheduleCreateSession2(session, validTimes))

	return nil
}

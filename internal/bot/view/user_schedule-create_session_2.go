package view

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/missingsemi/capstone/internal/bot/util"
	"github.com/missingsemi/capstone/internal/model"
	"github.com/slack-go/slack"
)

func UserScheduleCreateSession2(partialSession model.ScheduleSession, validTimes []time.Time) slack.ModalViewRequest {
	var modalRequest slack.ModalViewRequest

	timeSelectOptions := make([]*slack.OptionBlockObject, len(validTimes))

	for i, timeSlot := range validTimes {
		beginTime := timeSlot.Format(util.FriendlyFormat)
		timeSelectOptions[i] = slack.NewOptionBlockObject(
			timeSlot.Format(time.RFC3339),
			slack.NewTextBlockObject(
				"plain_text",
				fmt.Sprintf("%v", beginTime),
				false,
				false,
			),
			nil,
		)
	}

	modalRequest.Type = slack.ViewType("modal")
	modalRequest.Title = slack.NewTextBlockObject(
		"plain_text",
		"New Session",
		false,
		false,
	)
	modalRequest.Close = slack.NewTextBlockObject(
		"plain_text",
		"Cancel",
		false,
		false,
	)
	modalRequest.Submit = slack.NewTextBlockObject(
		"plain_text",
		"Submit",
		false,
		false,
	)
	modalRequest.Blocks = slack.Blocks{
		BlockSet: []slack.Block{
			slack.NewDividerBlock(),
			slack.NewInputBlock(
				"time_input_block",
				slack.NewTextBlockObject("plain_text", "Start Time", false, false),
				slack.NewOptionsSelectBlockElement(
					"static_select",
					slack.NewTextBlockObject("plain_text", "Select a start time", false, false),
					"time_input",
					timeSelectOptions...,
				),
			),
		},
	}
	modalRequest.CallbackID = "user_schedule-create_session_2-callback"

	sessionStr, _ := json.Marshal(partialSession)
	modalRequest.PrivateMetadata = string(sessionStr)

	return modalRequest
}

package view

import (
	"fmt"
	"time"

	"github.com/slack-go/slack"
)

func ScheduleAddTimeInformation(validTimes []time.Time) slack.ModalViewRequest {
	var modalRequest slack.ModalViewRequest

	timeSelectOptions := make([]*slack.OptionBlockObject, len(validTimes))

	for i, timeSlot := range validTimes {
		beginTime := timeSlot.Format(time.RFC3339)
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
		"Time Information",
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
				slack.NewTextBlockObject("plain_text", "Pick a time slot to reserve.", false, false),
				slack.NewOptionsSelectBlockElement(
					"static_select",
					slack.NewTextBlockObject("plain_text", "Select a time slot", false, false),
					"time_input",
					timeSelectOptions...,
				),
			),
		},
	}
	modalRequest.CallbackID = "schedule_add-time_information-callback"

	return modalRequest
}

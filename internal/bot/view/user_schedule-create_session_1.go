package view

import (
	"github.com/missingsemi/capstone/pkg/model"
	"github.com/slack-go/slack"
)

func UserScheduleCreateSession1(machines []model.Machine) slack.ModalViewRequest {
	var modalRequest slack.ModalViewRequest

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
		"Next",
		false,
		false,
	)

	groupInputBlock := slack.NewInputBlock(
		"group_input_block",
		slack.NewTextBlockObject(
			"plain_text",
			"Group Members",
			false,
			false,
		),
		slack.NewOptionsMultiSelectBlockElement(
			"multi_users_select",
			slack.NewTextBlockObject(
				"plain_text",
				"Add your group members (including yourself)",
				false,
				false,
			),
			"group_input",
		),
	)

	machineSelectOptions := make([]*slack.OptionBlockObject, len(machines))
	for i, machine := range machines {
		machineSelectOptions[i] = slack.NewOptionBlockObject(
			machine.Id,
			slack.NewTextBlockObject(
				"plain_text",
				machine.TitleName,
				false,
				false,
			),
			nil,
		)
	}

	machineInputBlock := slack.NewInputBlock(
		"machine_input_block",
		slack.NewTextBlockObject(
			"plain_text",
			"Machine",
			false,
			false,
		),
		slack.NewOptionsSelectBlockElement(
			"static_select",
			slack.NewTextBlockObject(
				"plain_text",
				"Select a machine",
				false,
				false,
			),
			"machine_input",
			machineSelectOptions...,
		),
	)

	reasonInputBlock := slack.NewInputBlock(
		"reason_input_block",
		slack.NewTextBlockObject(
			"plain_text",
			"Reason",
			false,
			false,
		),
		slack.NewPlainTextInputBlockElement(
			slack.NewTextBlockObject(
				"plain_text",
				"Write why you want to use this machine",
				false,
				false,
			),
			"reason_input",
		),
	)

	durationInputBlock := slack.NewInputBlock(
		"duration_input_block",
		slack.NewTextBlockObject(
			"plain_text",
			"Duration",
			false,
			false,
		),
		slack.NewOptionsSelectBlockElement(
			"static_select",
			slack.NewTextBlockObject(
				"plain_text",
				"Select how long you will need the machine for",
				false,
				false,
			),
			"duration_input",
			slack.NewOptionBlockObject(
				"15",
				slack.NewTextBlockObject(
					"plain_text",
					"15 Minutes",
					false,
					false,
				),
				nil,
			),
			slack.NewOptionBlockObject(
				"30",
				slack.NewTextBlockObject(
					"plain_text",
					"30 Minutes",
					false,
					false,
				),
				nil,
			),
			slack.NewOptionBlockObject(
				"45",
				slack.NewTextBlockObject(
					"plain_text",
					"45 Minutes",
					false,
					false,
				),
				nil,
			),
			slack.NewOptionBlockObject(
				"60",
				slack.NewTextBlockObject(
					"plain_text",
					"1 Hour",
					false,
					false,
				),
				nil,
			),
			slack.NewOptionBlockObject(
				"120",
				slack.NewTextBlockObject(
					"plain_text",
					"2 Hours",
					false,
					false,
				),
				nil,
			),
			slack.NewOptionBlockObject(
				"180",
				slack.NewTextBlockObject(
					"plain_text",
					"3 Hours",
					false,
					false,
				),
				nil,
			),
			slack.NewOptionBlockObject(
				"240",
				slack.NewTextBlockObject(
					"plain_text",
					"4 Hours",
					false,
					false,
				),
				nil,
			),
		),
	)

	modalRequest.Blocks = slack.Blocks{
		BlockSet: []slack.Block{
			slack.NewDividerBlock(),
			groupInputBlock,
			machineInputBlock,
			reasonInputBlock,
			durationInputBlock,
		},
	}

	modalRequest.CallbackID = "user_schedule-create_session_1-callback"

	return modalRequest
}

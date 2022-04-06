package view

import (
	"github.com/missingsemi/capstone/model"
	"github.com/slack-go/slack"
)

func ScheduleAddTeamInformation(machines []model.Machine) slack.ModalViewRequest {
	var modalRequest slack.ModalViewRequest

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

	modalRequest.Type = slack.ViewType("modal")
	modalRequest.Title = slack.NewTextBlockObject(
		"plain_text",
		"Team Information",
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
	modalRequest.Blocks = slack.Blocks{
		BlockSet: []slack.Block{
			slack.NewDividerBlock(),
			slack.NewInputBlock(
				"teammates_input_block",
				slack.NewTextBlockObject("plain_text", "Select the members of your group.", false, false),
				slack.NewOptionsMultiSelectBlockElement(
					"multi_users_select",
					slack.NewTextBlockObject(
						"plain_text",
						"Select users",
						false,
						false,
					),
					"teammates_input",
				),
			),
			slack.NewInputBlock(
				"machine_input_block",
				slack.NewTextBlockObject("plain_text", "Select the machine you want to reserve.", false, false),
				slack.NewOptionsSelectBlockElement(
					"static_select",
					slack.NewTextBlockObject("plain_text", "Select a machine", false, false),
					"machine_input",
					machineSelectOptions...,
				),
			),
		},
	}
	modalRequest.CallbackID = "schedule_add-team_information-callback"

	return modalRequest
}

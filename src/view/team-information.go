package view

import (
	"github.com/missingsemi/capstone/model"
	"github.com/slack-go/slack"
)

func TeamInformation() slack.ModalViewRequest {
	var modalRequest slack.ModalViewRequest

	machineSelectOptions := make([]*slack.OptionBlockObject, len(model.Machines))
	for i, machine := range model.Machines {
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
				"users_input_block",
				slack.NewTextBlockObject("plain_text", "Select the members of your group.", false, false),
				slack.NewOptionsMultiSelectBlockElement(
					"multi_users_select",
					slack.NewTextBlockObject(
						"plain_text",
						"Select users",
						false,
						false,
					),
					"users_select",
				),
			),
			slack.NewInputBlock(
				"machine_input_block",
				slack.NewTextBlockObject("plain_text", "Select the machine you want to reserve.", false, false),
				slack.NewOptionsSelectBlockElement(
					"static_select",
					slack.NewTextBlockObject("plain_text", "Select a machine", false, false),
					"machine_select",
					machineSelectOptions...,
				),
			),
			slack.NewDividerBlock(),
		},
	}
	modalRequest.CallbackID = "team-information-callback"

	return modalRequest
}

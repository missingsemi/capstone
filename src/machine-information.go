package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

func MachineInformation(session *SessionInfo) slack.ModalViewRequest {
	var modalRequest slack.ModalViewRequest

	machine := MachineFromId(session.Machine)

	title := machine.TitleName
	purposeQuestion := fmt.Sprintf("What do you want to use the %s for?", machine.Name)

	modalRequest.Type = slack.ViewType("modal")
	modalRequest.Title = slack.NewTextBlockObject(
		"plain_text",
		title,
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
				"machine_purpose_input_block",
				slack.NewTextBlockObject(
					"plain_text",
					purposeQuestion,
					false,
					false,
				),
				slack.NewPlainTextInputBlockElement(
					slack.NewTextBlockObject(
						"plain_text",
						"Purpose",
						false,
						false,
					),
					"machine_purpose",
				),
			),
			slack.NewInputBlock(
				"duration_input_block",
				slack.NewTextBlockObject(
					"plain_text",
					"How long will the part take to machine?",
					false,
					false,
				),
				slack.NewOptionsSelectBlockElement(
					"static_select",
					slack.NewTextBlockObject("plain_text", "Select a duration", false, false),
					"duration_select",
					slack.NewOptionBlockObject(
						"five_min",
						slack.NewTextBlockObject(
							"plain_text",
							"5 Minutes",
							false,
							false,
						),
						nil,
					),
					slack.NewOptionBlockObject(
						"ten_min",
						slack.NewTextBlockObject(
							"plain_text",
							"10 Minutes",
							false,
							false,
						),
						nil,
					),
					slack.NewOptionBlockObject(
						"fifteen_min",
						slack.NewTextBlockObject(
							"plain_text",
							"15 Minutes",
							false,
							false,
						),
						nil,
					),
					slack.NewOptionBlockObject(
						"thirty_min",
						slack.NewTextBlockObject(
							"plain_text",
							"30 Minutes",
							false,
							false,
						),
						nil,
					),
					slack.NewOptionBlockObject(
						"fortyfive_min",
						slack.NewTextBlockObject(
							"plain_text",
							"45 Minutes",
							false,
							false,
						),
						nil,
					),
					slack.NewOptionBlockObject(
						"one_hr",
						slack.NewTextBlockObject(
							"plain_text",
							"1 Hour",
							false,
							false,
						),
						nil,
					),
					slack.NewOptionBlockObject(
						"two_hr",
						slack.NewTextBlockObject(
							"plain_text",
							"2 Hours",
							false,
							false,
						),
						nil,
					),
					slack.NewOptionBlockObject(
						"three_hr",
						slack.NewTextBlockObject(
							"plain_text",
							"3 Hours",
							false,
							false,
						),
						nil,
					),
					slack.NewOptionBlockObject(
						"four_hr",
						slack.NewTextBlockObject(
							"plain_text",
							"4 Hours",
							false,
							false,
						),
						nil,
					),
				),
			),
			slack.NewInputBlock(
				"file_ready_input_block",
				slack.NewTextBlockObject(
					"plain_text",
					"Are the necessary files ready?",
					false,
					false,
				),
				slack.NewRadioButtonsBlockElement(
					"file_ready",
					slack.NewOptionBlockObject(
						"file_is_ready",
						slack.NewTextBlockObject(
							"plain_text",
							"Yes",
							false,
							false,
						),
						nil,
					),
				),
			),
		},
	}

	modalRequest.CallbackID = "machine-information-callback"

	return modalRequest
}

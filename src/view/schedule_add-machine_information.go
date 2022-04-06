package view

import (
	"fmt"

	"github.com/missingsemi/capstone/model"
	"github.com/slack-go/slack"
)

func ScheduleAddMachineInformation(machine model.Machine) slack.ModalViewRequest {
	var modalRequest slack.ModalViewRequest

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
				"reason_input_block",
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
					"reason_input",
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
					"file_ready_input",
					slack.NewOptionBlockObject(
						"file_ready_option",
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

	modalRequest.CallbackID = "schedule_add-machine_information-callback"

	return modalRequest
}

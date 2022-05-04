package view

import (
	"github.com/slack-go/slack"
)

func AdminMachinesCreateMachine() slack.ModalViewRequest {
	var modalRequest slack.ModalViewRequest

	modalRequest.Type = slack.ViewType("modal")
	modalRequest.Title = slack.NewTextBlockObject(
		"plain_text",
		"Create Machine",
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

	idInputBlock := slack.NewInputBlock(
		"id_input_block",
		slack.NewTextBlockObject(
			"plain_text",
			"Id",
			false,
			false,
		),
		slack.NewPlainTextInputBlockElement(
			slack.NewTextBlockObject(
				"plain_text",
				"machine_id",
				false,
				false,
			),
			"id_input",
		),
	)
	idInputBlock.Hint = slack.NewTextBlockObject(
		"plain_text",
		"The id of the machine, must be unique.",
		false,
		false,
	)

	nameInputBlock := slack.NewInputBlock(
		"name_input_block",
		slack.NewTextBlockObject(
			"plain_text",
			"Name",
			false,
			false,
		),
		slack.NewPlainTextInputBlockElement(
			slack.NewTextBlockObject(
				"plain_text",
				"machine name",
				false,
				false,
			),
			"name_input",
		),
	)
	nameInputBlock.Hint = slack.NewTextBlockObject(
		"plain_text",
		"The lowercase name of the machine for use in sentences.",
		false,
		false,
	)

	titleNameInputBlock := slack.NewInputBlock(
		"titlename_input_block",
		slack.NewTextBlockObject(
			"plain_text",
			"Title Name",
			false,
			false,
		),
		slack.NewPlainTextInputBlockElement(
			slack.NewTextBlockObject(
				"plain_text",
				"Machine Title Name",
				false,
				false,
			),
			"titlename_input",
		),
	)
	titleNameInputBlock.Hint = slack.NewTextBlockObject(
		"plain_text",
		"The titlecase name of the machine for use in titles and other labels.",
		false,
		false,
	)

	countInputBlock := slack.NewInputBlock(
		"count_input_block",
		slack.NewTextBlockObject(
			"plain_text",
			"Count",
			false,
			false,
		),
		slack.NewPlainTextInputBlockElement(
			slack.NewTextBlockObject(
				"plain_text",
				"0",
				false,
				false,
			),
			"count_input",
		),
	)
	countInputBlock.Hint = slack.NewTextBlockObject(
		"plain_text",
		"The number of machines of this type.",
		false,
		false,
	)

	modalRequest.Blocks = slack.Blocks{
		BlockSet: []slack.Block{
			slack.NewDividerBlock(),
			idInputBlock,
			nameInputBlock,
			titleNameInputBlock,
			countInputBlock,
		},
	}

	modalRequest.CallbackID = "admin_machines-create_machine-callback"

	return modalRequest
}

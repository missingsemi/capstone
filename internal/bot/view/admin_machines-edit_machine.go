package view

import (
	"fmt"

	"github.com/missingsemi/capstone/pkg/model"
	"github.com/slack-go/slack"
)

func AdminMachinesEditMachine(machine model.Machine) slack.ModalViewRequest {
	var modalRequest slack.ModalViewRequest

	modalRequest.Type = slack.ViewType("modal")
	modalRequest.Title = slack.NewTextBlockObject(
		"plain_text",
		fmt.Sprintf("Edit Machine"),
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
				machine.Id,
				false,
				false,
			),
			"id_input",
		),
	)
	idInputBlock.Optional = true
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
				machine.Name,
				false,
				false,
			),
			"name_input",
		),
	)
	nameInputBlock.Optional = true
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
				machine.TitleName,
				false,
				false,
			),
			"titlename_input",
		),
	)
	titleNameInputBlock.Optional = true
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
				fmt.Sprintf("%d", machine.Count),
				false,
				false,
			),
			"count_input",
		),
	)
	countInputBlock.Optional = true
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

	modalRequest.CallbackID = "admin_machines-edit_machine-callback"

	modalRequest.PrivateMetadata = machine.Id

	return modalRequest
}

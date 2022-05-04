package view

import (
	"fmt"

	"github.com/missingsemi/capstone/pkg/model"
	"github.com/slack-go/slack"
)

func AdminMachinesAvailableMachines(machines []model.Machine) slack.ModalViewRequest {
	var modalRequest slack.ModalViewRequest

	machineBlocks := make([]slack.Block, 3*len(machines))
	for i, machine := range machines {
		editButton := slack.NewButtonBlockElement(
			"admin_machines-available_machines-edit_callback",
			machine.Id,
			slack.NewTextBlockObject(
				"plain_text",
				"Edit",
				false,
				false,
			),
		)

		deleteButton := slack.NewButtonBlockElement(
			"admin_machines-available_machines-delete_callback",
			machine.Id,
			slack.NewTextBlockObject(
				"plain_text",
				"Delete",
				false,
				false,
			),
		)
		deleteButton.Style = "danger"
		deleteButton.Confirm = slack.NewConfirmationBlockObject(
			slack.NewTextBlockObject(
				"plain_text",
				fmt.Sprintf("Are you sure you want to delete %s?", machine.TitleName),
				false,
				false,
			),
			slack.NewTextBlockObject(
				"mrkdwn",
				"Deleting this machine will also delete every schedule for the machine. This action *cannot* be undone.",
				false,
				false,
			),
			slack.NewTextBlockObject(
				"plain_text",
				"Delete",
				false,
				false,
			),
			slack.NewTextBlockObject(
				"plain_text",
				"Cancel",
				false,
				false,
			),
		)
		deleteButton.Confirm.Style = "danger"

		machineBlocks[3*i] = slack.NewDividerBlock()

		machineBlocks[3*i+1] = slack.NewSectionBlock(
			slack.NewTextBlockObject(
				"mrkdwn",
				fmt.Sprintf("*%s* - %d available\n%s", machine.TitleName, machine.Count, machine.Id),
				false,
				false,
			),
			nil,
			nil,
		)

		machineBlocks[3*i+2] = slack.NewActionBlock(
			"",
			editButton,
			deleteButton,
		)
	}

	modalRequest.Type = slack.ViewType("modal")
	modalRequest.Title = slack.NewTextBlockObject(
		"plain_text",
		"Machine List",
		false,
		false,
	)
	modalRequest.Close = slack.NewTextBlockObject(
		"plain_text",
		"Close",
		false,
		false,
	)

	createMachineButton := slack.NewButtonBlockElement(
		"admin_machines-available_machines-create_callback",
		"",
		slack.NewTextBlockObject(
			"plain_text",
			"Add a Machine",
			false,
			false,
		),
	)
	createMachineButton.Style = "primary"

	summary := slack.NewSectionBlock(
		slack.NewTextBlockObject(
			"plain_text",
			fmt.Sprintf("There are currently %d types of machine.", len(machines)),
			false,
			false,
		),
		nil,
		slack.NewAccessory(createMachineButton),
	)
	modalRequest.Blocks = slack.Blocks{
		BlockSet: append(
			[]slack.Block{summary},
			machineBlocks...,
		),
	}
	modalRequest.CallbackID = "admin_machines-available_machines-callback"

	return modalRequest
}

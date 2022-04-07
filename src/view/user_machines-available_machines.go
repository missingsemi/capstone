package view

import (
	"fmt"

	"github.com/missingsemi/capstone/model"
	"github.com/slack-go/slack"
)

func UserMachinesAvailableMachines(machines []model.Machine) slack.ModalViewRequest {
	var modalRequest slack.ModalViewRequest

	machineBlocks := make([]slack.Block, 2*len(machines))
	for i, machine := range machines {
		machineBlocks[2*i] = slack.NewDividerBlock()

		machineBlocks[2*i+1] = slack.NewSectionBlock(
			slack.NewTextBlockObject(
				"mrkdwn",
				fmt.Sprintf("*%s* - %d available", machine.TitleName, machine.Count),
				false,
				false,
			),
			nil,
			nil,
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

	summary := slack.NewSectionBlock(
		slack.NewTextBlockObject(
			"plain_text",
			fmt.Sprintf("There are currently %d types of machine.", len(machines)),
			false,
			false,
		),
		nil,
		nil,
	)
	modalRequest.Blocks = slack.Blocks{
		BlockSet: append(
			[]slack.Block{summary},
			machineBlocks...,
		),
	}

	return modalRequest
}

package view

import (
	"fmt"
	"time"

	"github.com/missingsemi/capstone/model"
	"github.com/missingsemi/capstone/util"
	"github.com/slack-go/slack"
)

func UserScheduleCreatedSessions(sessions []model.ScheduleSession, machines []model.Machine) slack.ModalViewRequest {
	var modalRequest slack.ModalViewRequest

	sessionBlocks := make([]slack.Block, 3*len(sessions))
	for i, session := range sessions {
		machine := util.FilterMachine(session.Machine, machines)

		editButton := slack.NewButtonBlockElement(
			"user_schedule-created_sessions-view_callback",
			fmt.Sprintf("%d", session.Id),
			slack.NewTextBlockObject(
				"plain_text",
				"View Details",
				false,
				false,
			),
		)

		deleteButton := slack.NewButtonBlockElement(
			"user_schedule-created_sessions-delete_callback",
			fmt.Sprintf("%d", session.Id),
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
				"Are you sure you want to delete this session?",
				false,
				false,
			),
			slack.NewTextBlockObject(
				"mrkdwn",
				"If you delete this session, you may not be able to get it back later. This action *cannot* be undone.",
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

		sessionBlocks[3*i] = slack.NewDividerBlock()

		sessionBlocks[3*i+1] = slack.NewSectionBlock(
			slack.NewTextBlockObject(
				"mrkdwn",
				fmt.Sprintf("*%s* @ %s - %d Minutes", machine.TitleName, session.Time.Format(time.RFC3339), session.Duration),
				false,
				false,
			),
			nil,
			nil,
		)

		sessionBlocks[3*i+2] = slack.NewActionBlock(
			"",
			editButton,
			deleteButton,
		)
	}

	modalRequest.Type = slack.ViewType("modal")
	modalRequest.Title = slack.NewTextBlockObject(
		"plain_text",
		"Your Schedule",
		false,
		false,
	)
	modalRequest.Close = slack.NewTextBlockObject(
		"plain_text",
		"Close",
		false,
		false,
	)

	createSessionButton := slack.NewButtonBlockElement(
		"user_schedule-created_sessions-create_callback",
		"",
		slack.NewTextBlockObject(
			"plain_text",
			"New Session",
			false,
			false,
		),
	)
	createSessionButton.Style = "primary"

	summary := slack.NewSectionBlock(
		slack.NewTextBlockObject(
			"plain_text",
			fmt.Sprintf("You currently have %d upcoming session(s) scheduled.", len(sessions)),
			false,
			false,
		),
		nil,
		slack.NewAccessory(createSessionButton),
	)
	modalRequest.Blocks = slack.Blocks{
		BlockSet: append(
			[]slack.Block{summary},
			sessionBlocks...,
		),
	}
	modalRequest.CallbackID = "user_schedule-created_sessions-callback"

	return modalRequest
}

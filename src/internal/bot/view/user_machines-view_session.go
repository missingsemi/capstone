package view

import (
	"fmt"
	"strings"

	"github.com/missingsemi/capstone/internal/bot/util"
	"github.com/missingsemi/capstone/pkg/model"
	"github.com/slack-go/slack"
)

func UserMachinesViewSession(session model.ScheduleSession, machine model.Machine) slack.ModalViewRequest {
	var modalRequest slack.ModalViewRequest

	modalRequest.Type = slack.ViewType("modal")
	modalRequest.Title = slack.NewTextBlockObject(
		"plain_text",
		fmt.Sprintf("View Session"),
		false,
		false,
	)
	modalRequest.Close = slack.NewTextBlockObject(
		"plain_text",
		"Cancel",
		false,
		false,
	)

	mainSection := slack.NewSectionBlock(
		nil,
		[]*slack.TextBlockObject{
			slack.NewTextBlockObject(
				"mrkdwn",
				fmt.Sprintf("*Machine*\n%s", machine.TitleName),
				false,
				false,
			),
			slack.NewTextBlockObject(
				"mrkdwn",
				fmt.Sprintf("*Reason*\n%s", session.Reason),
				false,
				false,
			),
			slack.NewTextBlockObject(
				"mrkdwn",
				fmt.Sprintf("*Start Time*\n%s", session.Time.Format(util.FriendlyFormat)),
				false,
				false,
			),
			slack.NewTextBlockObject(
				"mrkdwn",
				fmt.Sprintf("*Duration*\n%d Minutes", session.Duration),
				false,
				false,
			),
		},
		nil,
	)

	groupMemberString := "<@" + strings.Join(session.GroupIds, "> <@") + ">"

	groupMemberSection := slack.NewSectionBlock(
		slack.NewTextBlockObject(
			"mrkdwn",
			"*Group Members*\n"+groupMemberString,
			false,
			false,
		),
		nil,
		nil,
	)

	modalRequest.Blocks = slack.Blocks{
		BlockSet: []slack.Block{
			slack.NewDividerBlock(),
			mainSection,
			groupMemberSection,
		},
	}

	return modalRequest
}

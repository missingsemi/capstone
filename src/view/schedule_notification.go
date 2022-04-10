package view

import (
	"fmt"

	"github.com/missingsemi/capstone/model"
	"github.com/missingsemi/capstone/util"
	"github.com/slack-go/slack"
)

func ScheduleNotification(session model.ScheduleSession, machine model.Machine) []slack.MsgOption {
	if session.Stage == 0 {
		return []slack.MsgOption{
			slack.MsgOptionText(
				fmt.Sprintf("Your %s session for the %s will be starting in less than five minutes.", session.Time.Format(util.FriendlyFormat), machine.Name),
				false,
			),
		}
	}

	if session.Stage == 1 {
		return []slack.MsgOption{
			slack.MsgOptionText(
				fmt.Sprintf("Your session for the %s has now started, and it will end in %d minutes.", machine.Name, session.Duration),
				false,
			),
		}
	}

	if session.Stage == 2 {
		return []slack.MsgOption{
			slack.MsgOptionText(
				fmt.Sprintf("Your session for the %s will be ending in less than five minutes. Try to have your part removed and the machine cleaned soon.", machine.Name),
				false,
			),
		}
	}

	if session.Stage == 3 {
		return []slack.MsgOption{
			slack.MsgOptionText(
				fmt.Sprintf("Your session for the %s has ended.", machine.Name),
				false,
			),
		}
	}

	return []slack.MsgOption{}
}

package util

import (
	"time"

	"github.com/missingsemi/capstone/database"
	"github.com/missingsemi/capstone/model"
	"github.com/missingsemi/capstone/view"
	"github.com/slack-go/slack"
)

func Notify(api *slack.Client) {
	for range time.Tick(time.Minute) {
		sessions, err := database.GetUnfinishedSessions()
		if err != nil {
			return
		}
		machines, err := database.GetMachines()
		if err != nil {
			return
		}

		for _, session := range sessions {
			machine := filterMachine(session.Machine, machines)

			if session.Stage == 0 && session.Time.Sub(time.Now()) < 5*time.Minute {
				// notify that session is coming up in < 5 mins
				msg := view.ScheduleNotification(session, machine)
				api.PostMessage(session.UserId, msg...)
				session.Stage = 1
				database.ModifySession(session.Id, session)
				continue
			}

			if session.Stage == 1 && session.Time.Before(time.Now()) {
				// notify that session has started
				msg := view.ScheduleNotification(session, machine)
				api.PostMessage(session.UserId, msg...)
				session.Stage = 2
				database.ModifySession(session.Id, session)
				continue
			}

			endTime := session.Time.Add(time.Duration(session.Duration) * time.Minute)
			if session.Stage == 2 && endTime.Sub(time.Now()) < 5*time.Minute {
				msg := view.ScheduleNotification(session, machine)
				api.PostMessage(session.UserId, msg...)
				session.Stage = 3
				database.ModifySession(session.Id, session)
				continue
			}

			if session.Stage == 3 && endTime.Before(time.Now()) {
				msg := view.ScheduleNotification(session, machine)
				api.PostMessage(session.UserId, msg...)
				session.Stage = 4
				database.ModifySession(session.Id, session)
			}
		}
	}
}

func filterMachine(id string, machines []model.Machine) model.Machine {
	for _, machine := range machines {
		if machine.Id == id {
			return machine
		}
	}

	return model.Machine{}
}

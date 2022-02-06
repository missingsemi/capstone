package main

import (
	"fmt"
	"time"

	"github.com/slack-go/slack"
)

func TimeInformation(session *SessionInfo, sessions []*SessionInfo) slack.ModalViewRequest {
	var modalRequest slack.ModalViewRequest
	duration := time.Duration(1000000000 * 60 * session.Duration)

	if len(sessions) > 5 {
		sessions = sessions[0:5]
	}

	validTimes := ValidTimes(session.Machine, session.Duration, sessions)
	timeSelectOptions := make([]*slack.OptionBlockObject, len(validTimes))
	for i, timeSlot := range validTimes {
		beginTime := timeSlot.Format(time.Stamp)
		endTime := timeSlot.Add(duration).Format(time.Stamp)
		timeSelectOptions[i] = slack.NewOptionBlockObject(
			timeSlot.Format(time.RFC3339),
			slack.NewTextBlockObject(
				"plain_text",
				fmt.Sprintf("%s to %s", beginTime, endTime),
				false,
				false,
			),
			nil,
		)
	}

	modalRequest.Type = slack.ViewType("modal")
	modalRequest.Title = slack.NewTextBlockObject(
		"plain_text",
		"Time Information",
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
	modalRequest.Blocks = slack.Blocks{
		BlockSet: []slack.Block{
			slack.NewDividerBlock(),
			slack.NewInputBlock(
				"time_input_block",
				slack.NewTextBlockObject("plain_text", "Pick a time slot to reserve.", false, false),
				slack.NewOptionsSelectBlockElement(
					"static_select",
					slack.NewTextBlockObject("plain_text", "Select a time slot", false, false),
					"time_select",
					timeSelectOptions...,
				),
			),
			slack.NewDividerBlock(),
		},
	}
	modalRequest.CallbackID = "time-information-callback"

	return modalRequest
}

func ValidTimes(machine string, duration int64, sessions []*SessionInfo) []time.Time {
	machineInfo := MachineFromId(machine)
	if machineInfo == nil {
		return []time.Time{}
	}

	filtered := FilterSessions(sessions, FilterOptions{
		Machine: machine,
	})

	// Granularity of five minutes, 300 sec
	slots := make(map[int64]int)
	for _, session := range filtered {
		limit := session.Time.UTC().Unix() + session.Duration*60
		for t := session.Time.UTC().Unix(); t <= limit; t += 300 {
			slots[t] += 1
		}
	}

	blockLength := int64(0)
	validBlocks := make([]time.Time, 0)
	now := (time.Now().Unix()/300 + 1) * 300 // Round up to nearest 5 minute block
	for t := now; t < now+24*60*60; t += 300 {
		//fmt.Println("slots[", t, "] =", slots[t])
		// If its during a valid timeframe and the machine isn't booked,
		if DuringBreak(t) {
			blockLength = 0
			continue
		}
		if slots[t] < machineInfo.Count {
			blockLength++
		} else {
			blockLength = 0
		}
		// and if the block of time is long enough
		if blockLength*5 >= duration {
			// add the start of this time frame as a valid option
			// fmt.Println("Found a block of length", blockLength, "at time", t-duration*60)
			validBlocks = append(
				validBlocks,
				time.Unix(t-duration*60, 0),
			)

			if len(validBlocks) == 5 {
				break
			}
		}
	}

	fmt.Println("BLOCKS: ", validBlocks)

	return validBlocks
}

func DuringBreak(t int64) bool {

	return false

	timeObj := time.Unix(t, 0)
	// don't let people schedule machines after 5pm or before 8am
	// TODO: Config this and make it changeable in admin command
	if timeObj.Hour() > 17 || timeObj.Hour() < 8 {
		return true
	}
	if timeObj.Weekday() == time.Sunday || timeObj.Weekday() == time.Saturday {
		return true
	}

	return false
}

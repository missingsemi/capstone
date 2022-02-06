package main

import (
	"strconv"
	"strings"
	"time"
)

type SessionInfo struct {
	ViewId   string
	UserId   string
	UserName string
	GroupIds []string
	Machine  string
	Reason   string
	Duration int64
	Time     time.Time
}

func SessionFromRow(row []interface{}) *SessionInfo {
	userId, ok := row[0].(string)
	if !ok {
		return nil
	}

	userName, ok := row[1].(string)
	if !ok {
		return nil
	}

	groupIdsString, ok := row[2].(string)
	if !ok {
		return nil
	}
	groupIds := strings.Split(groupIdsString, ",")

	machine, ok := row[3].(string)
	if !ok {
		return nil
	}

	reason, ok := row[4].(string)
	if !ok {
		return nil
	}

	durationString, ok := row[5].(string)
	if !ok {
		return nil
	}
	duration, err := strconv.ParseInt(durationString, 10, 64)
	if err != nil {
		return nil
	}

	timeString, ok := row[6].(string)
	if !ok {
		return nil
	}
	timeObj, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		return nil
	}

	return &SessionInfo{
		UserId:   userId,
		UserName: userName,
		GroupIds: groupIds,
		Machine:  machine,
		Reason:   reason,
		Duration: duration,
		Time:     timeObj,
	}
}

func SessionToRow(s *SessionInfo) []interface{} {
	result := make([]interface{}, 7)
	result[0] = s.UserId
	result[1] = s.UserName
	result[2] = strings.Join(s.GroupIds, ",")
	result[3] = s.Machine
	result[4] = s.Reason
	result[5] = s.Duration
	result[6] = s.Time.Format(time.RFC3339)

	return result
}

func DurationToMins(duration string) int64 {
	switch duration {
	case "five_min":
		return 5
	case "ten_min":
		return 10
	case "fifteen_min":
		return 15
	case "thirty_min":
		return 30
	case "fortyfive_min":
		return 45
	case "one_hr":
		return 60
	case "two_hr":
		return 120
	case "three_hr":
		return 180
	case "four_hr":
		return 240
	default:
		return 0
	}
}

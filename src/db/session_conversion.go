package db

import (
	"strconv"
	"strings"
	"time"

	"github.com/missingsemi/capstone/model"
)

func SessionFromRow(row []interface{}) *model.ScheduleAddSession {
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

	return &model.ScheduleAddSession{
		UserId:   userId,
		UserName: userName,
		GroupIds: groupIds,
		Machine:  machine,
		Reason:   reason,
		Duration: duration,
		Time:     timeObj,
	}
}

func SessionToRow(s *model.ScheduleAddSession) []interface{} {
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

package model

import (
	"time"
)

type ScheduleAddSession struct {
	ViewId   string
	UserId   string
	UserName string
	GroupIds []string
	Machine  string
	Reason   string
	Duration int64
	Time     time.Time
}

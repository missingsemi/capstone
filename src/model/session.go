package model

import (
	"time"
)

type ScheduleSession struct {
	Id       int
	UserId   string
	GroupIds []string
	Machine  string
	Reason   string
	Duration int64
	Time     time.Time
	Stage    int
}

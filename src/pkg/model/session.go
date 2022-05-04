package model

import (
	"time"
)

type ScheduleSession struct {
	Id       int       `json:"id"`
	UserId   string    `json:"userId"`
	GroupIds []string  `json:"groupIds"`
	Machine  string    `json:"machine"`
	Reason   string    `json:"reason"`
	Duration int64     `json:"duration"`
	Time     time.Time `json:"time"`
	Stage    int       `json:"-"`
}

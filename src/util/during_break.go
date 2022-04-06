package util

import "time"

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

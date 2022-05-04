package util

import "time"

// Assumes office hours are 8AM-5PM on weekdays.
func DuringOfficeHours(t time.Time) bool {

	// 17th hour goes 5pm-6pm, so hour should be 16 or less to go up until 5pm.
	if t.Hour() < 8 || t.Hour() > 16 {
		return false
	}

	if t.Weekday() == time.Sunday || t.Weekday() == time.Saturday {
		return false
	}

	return true
}

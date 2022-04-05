package util

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

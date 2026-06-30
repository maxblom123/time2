package time2

import (
	"fmt"
	"time"
)

const (
	Second = time.Second
	Minute = time.Minute
	Hour   = time.Hour
	Day    = 24 * Hour
	Week   = 7 * Day
	Month  = 30 * Day
	Year   = 365 * Day
)

func HumanizeDuration(d time.Duration) string {
	if d < 0 {
		return "0 seconds"
	}
	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	mins := int(d.Minutes()) % 60
	secs := int(d.Seconds()) % 60

	switch {
	case days > 0 && hours > 0:
		return pluralize(days, "day") + " " + pluralize(hours, "hour")
	case days > 0:
		return pluralize(days, "day")
	case hours > 0 && mins > 0:
		return pluralize(hours, "hour") + " " + pluralize(mins, "minute")
	case hours > 0:
		return pluralize(hours, "hour")
	case mins > 0:
		return pluralize(mins, "minute")
	default:
		return pluralize(secs, "second")
	}
}

func pluralize(n int, unit string) string {
	if n == 1 {
		return "1 " + unit
	}
	return fmt.Sprintf("%d %ss", n, unit)
}
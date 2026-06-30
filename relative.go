package time2

import "time"

func (t T) Humanize() string {
	diff := time.Since(t.Time)
	future := diff < 0
	if future {
		diff = -diff
	}

	var label string
	switch {
	case diff < 5*time.Second:
		return "just now"
	case diff < time.Minute:
		label = pluralize(int(diff.Seconds()), "second")
	case diff < time.Hour:
		label = pluralize(int(diff.Minutes()), "minute")
	case diff < Day:
		label = pluralize(int(diff.Hours()), "hour")
	case diff < Week:
		label = pluralize(int(diff.Hours()/24), "day")
	case diff < Month:
		label = pluralize(int(diff.Hours()/24/7), "week")
	case diff < Year:
		label = pluralize(int(diff.Hours()/24/30), "month")
	default:
		label = pluralize(int(diff.Hours()/24/365), "year")
	}

	if future {
		return "in " + label
	}
	return label + " ago"
}

func (t T) FromNow() string { return t.Humanize() }
func (t T) DaysUntil() int  { return int(time.Until(t.Time).Hours() / 24) }
func (t T) DaysSince() int  { return -t.DaysUntil() }
func (t T) Diff(u T) string { return HumanizeDuration(t.Time.Sub(u.Time)) }
package time2

import (
	"strings"
	"time"
)

func (t T) Format(layout string) string {
	r := strings.NewReplacer(
		"YYYY", t.Time.Format("2006"),
		"YY", t.Time.Format("06"),
		"MMMM", t.Time.Format("January"),
		"MMM", t.Time.Format("Jan"),
		"MM", t.Time.Format("01"),
		"DD", t.Time.Format("02"),
		"HH", t.Time.Format("15"),
		"hh", t.Time.Format("03"),
		"mm", t.Time.Format("04"),
		"ss", t.Time.Format("05"),
		"A", t.Time.Format("PM"),
		"dddd", t.Time.Format("Monday"),
		"ddd", t.Time.Format("Mon"),
	)
	return r.Replace(layout)
}

func (t T) ToDateString() string     { return t.Format("YYYY-MM-DD") }
func (t T) ToTimeString() string     { return t.Format("HH:mm:ss") }
func (t T) ToDateTimeString() string { return t.Format("YYYY-MM-DD HH:mm:ss") }
func (t T) ToFriendly() string       { return t.Format("dddd, MMMM DD YYYY") }
func (t T) ToRFC3339() string        { return t.Time.Format(time.RFC3339) }
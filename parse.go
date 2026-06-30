package time2

import (
	"fmt"
	"time"
)

var knownLayouts = []string{
	"2006-01-02 15:04:05",
	"2006-01-02T15:04:05Z07:00",
	"2006-01-02",
	"01/02/2006",
	"02-01-2006",
	"January 2, 2006",
	"Jan 2, 2006",
	"2 January 2006",
	time.RFC3339,
	time.RFC822,
	time.RFC1123,
}

func Parse(s string) (T, error) {
	for _, layout := range knownLayouts {
		if t, err := time.Parse(layout, s); err == nil {
			return T{t}, nil
		}
	}
	return T{}, fmt.Errorf("time2: cannot parse %q", s)
}

func MustParse(s string) T {
	t, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return t
}
package time2

import "time"

type T struct {
	time.Time
}

func Now() T                                                             { return T{time.Now()} }
func New(t time.Time) T                                                  { return T{t} }
func From(year int, month time.Month, day int) T                         { return T{time.Date(year, month, day, 0, 0, 0, 0, time.Local)} }
func FromDateTime(year int, month time.Month, day, hour, min, sec int) T { return T{time.Date(year, month, day, hour, min, sec, 0, time.Local)} }

func (t T) IsPast() bool   { return t.Time.Before(time.Now()) }
func (t T) IsFuture() bool { return t.Time.After(time.Now()) }
func (t T) IsToday() bool {
	y1, m1, d1 := t.Date()
	y2, m2, d2 := time.Now().Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
func (t T) IsWeekend() bool { w := t.Weekday(); return w == time.Saturday || w == time.Sunday }
func (t T) IsWeekday() bool { return !t.IsWeekend() }

func (t T) AddDays(n int) T   { return T{t.Time.Add(Day * time.Duration(n))} }
func (t T) AddWeeks(n int) T  { return T{t.Time.Add(Week * time.Duration(n))} }
func (t T) AddMonths(n int) T { return T{t.Time.AddDate(0, n, 0)} }
func (t T) AddYears(n int) T  { return T{t.Time.AddDate(n, 0, 0)} }

func (t T) StartOfDay() T {
	y, m, d := t.Date()
	return T{time.Date(y, m, d, 0, 0, 0, 0, t.Location())}
}

func (t T) EndOfDay() T {
	y, m, d := t.Date()
	return T{time.Date(y, m, d, 23, 59, 59, 999999999, t.Location())}
}

func (t T) StartOfWeek() T {
	offset := int(t.Weekday())
	if offset == 0 {
		offset = 7
	}
	return t.AddDays(-(offset - 1)).StartOfDay()
}

func (t T) StartOfMonth() T {
	y, m, _ := t.Date()
	return T{time.Date(y, m, 1, 0, 0, 0, 0, t.Location())}
}

func (t T) EndOfMonth() T {
	return T{t.StartOfMonth().AddMonths(1).Time.Add(-time.Nanosecond)}
}

func (t T) StartOfYear() T {
	return T{time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location())}
}
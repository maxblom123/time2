package time2

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	if time.Since(Now().Time) > time.Second {
		t.Fatal("Now() drifts beyond 1s from current time")
	}
}

func TestNew(t *testing.T) {
	now := time.Now()
	if !New(now).Time.Equal(now) {
		t.Fatal("New() must wrap the provided time exactly")
	}
}

func TestFrom(t *testing.T) {
	d := From(2026, time.June, 30)
	if d.Year() != 2026 || d.Month() != time.June || d.Day() != 30 {
		t.Fatalf("From() = %v", d)
	}
}

func TestFromDateTime(t *testing.T) {
	d := FromDateTime(2026, time.June, 30, 12, 30, 45)
	if d.Hour() != 12 || d.Minute() != 30 || d.Second() != 45 {
		t.Fatalf("FromDateTime() = %v", d)
	}
}

func TestIsPast(t *testing.T) {
	if !New(time.Now().Add(-time.Hour)).IsPast() {
		t.Fatal("past time must satisfy IsPast()")
	}
	if New(time.Now().Add(time.Hour)).IsPast() {
		t.Fatal("future time must not satisfy IsPast()")
	}
}

func TestIsFuture(t *testing.T) {
	if !New(time.Now().Add(time.Hour)).IsFuture() {
		t.Fatal("future time must satisfy IsFuture()")
	}
	if New(time.Now().Add(-time.Hour)).IsFuture() {
		t.Fatal("past time must not satisfy IsFuture()")
	}
}

func TestIsToday(t *testing.T) {
	if !Now().IsToday() {
		t.Fatal("Now() must be today")
	}
	if New(time.Now().Add(-48 * time.Hour)).IsToday() {
		t.Fatal("two days ago must not be today")
	}
}

func TestIsWeekend(t *testing.T) {
	cases := []struct {
		date    T
		weekend bool
	}{
		{From(2026, time.June, 27), true},
		{From(2026, time.June, 28), true},
		{From(2026, time.June, 29), false},
	}
	for _, c := range cases {
		if c.date.IsWeekend() != c.weekend {
			t.Errorf("%s: IsWeekend() = %v, want %v", c.date.Weekday(), c.date.IsWeekend(), c.weekend)
		}
	}
}

func TestIsWeekday(t *testing.T) {
	if !From(2026, time.June, 29).IsWeekday() {
		t.Fatal("Monday must be a weekday")
	}
	if From(2026, time.June, 27).IsWeekday() {
		t.Fatal("Saturday must not be a weekday")
	}
}

func TestAddDays(t *testing.T) {
	if got := From(2026, time.June, 28).AddDays(2).Day(); got != 30 {
		t.Fatalf("AddDays(2) = day %d, want 30", got)
	}
}

func TestAddWeeks(t *testing.T) {
	if got := From(2026, time.June, 1).AddWeeks(2).Day(); got != 15 {
		t.Fatalf("AddWeeks(2) = day %d, want 15", got)
	}
}

func TestAddMonths(t *testing.T) {
	if got := From(2026, time.January, 31).AddMonths(1).Month(); got != time.March {
		t.Fatalf("AddMonths(1) from Jan 31 = %s, want March", got)
	}
}

func TestAddYears(t *testing.T) {
	if got := From(2026, time.June, 30).AddYears(1).Year(); got != 2027 {
		t.Fatalf("AddYears(1) = %d, want 2027", got)
	}
}

func TestStartOfDay(t *testing.T) {
	s := FromDateTime(2026, time.June, 30, 15, 45, 30).StartOfDay()
	if s.Hour() != 0 || s.Minute() != 0 || s.Second() != 0 {
		t.Fatalf("StartOfDay() = %v, want midnight", s)
	}
}

func TestEndOfDay(t *testing.T) {
	e := From(2026, time.June, 30).EndOfDay()
	if e.Hour() != 23 || e.Minute() != 59 || e.Second() != 59 {
		t.Fatalf("EndOfDay() = %v, want 23:59:59", e)
	}
}

func TestStartOfWeek(t *testing.T) {
	cases := []struct {
		in  T
		day time.Weekday
	}{
		{From(2026, time.July, 1), time.Monday},
		{From(2026, time.June, 28), time.Monday},
	}
	for _, c := range cases {
		if got := c.in.StartOfWeek().Weekday(); got != c.day {
			t.Errorf("StartOfWeek() from %s = %s, want %s", c.in.Weekday(), got, c.day)
		}
	}
}

func TestStartOfMonth(t *testing.T) {
	if got := From(2026, time.June, 30).StartOfMonth().Day(); got != 1 {
		t.Fatalf("StartOfMonth() = day %d, want 1", got)
	}
}

func TestEndOfMonth(t *testing.T) {
	if got := From(2026, time.June, 15).EndOfMonth().Day(); got != 30 {
		t.Fatalf("EndOfMonth() = day %d, want 30", got)
	}
}

func TestStartOfYear(t *testing.T) {
	s := From(2026, time.June, 30).StartOfYear()
	if s.Month() != time.January || s.Day() != 1 {
		t.Fatalf("StartOfYear() = %v, want Jan 1", s)
	}
}

func TestHumanizeDuration(t *testing.T) {
	cases := []struct {
		in   time.Duration
		want string
	}{
		{-5 * time.Second, "0 seconds"},
		{time.Second, "1 second"},
		{30 * time.Second, "30 seconds"},
		{45 * time.Minute, "45 minutes"},
		{90 * time.Minute, "1 hour 30 minutes"},
		{2 * time.Hour, "2 hours"},
		{25 * time.Hour, "1 day 1 hour"},
		{48 * time.Hour, "2 days"},
	}
	for _, c := range cases {
		if got := HumanizeDuration(c.in); got != c.want {
			t.Errorf("HumanizeDuration(%v) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFormat(t *testing.T) {
	d := From(2026, time.June, 30)
	cases := []struct {
		layout string
		want   string
	}{
		{"YYYY-MM-DD", "2026-06-30"},
		{"YY", "26"},
		{"MMMM", "June"},
		{"MMM", "Jun"},
	}
	for _, c := range cases {
		if got := d.Format(c.layout); got != c.want {
			t.Errorf("Format(%q) = %q, want %q", c.layout, got, c.want)
		}
	}
}

func TestFormatTime(t *testing.T) {
	if got := FromDateTime(2026, time.June, 30, 14, 5, 9).Format("HH:mm:ss"); got != "14:05:09" {
		t.Fatalf("Format(HH:mm:ss) = %q, want 14:05:09", got)
	}
}

func TestFormatAMPM(t *testing.T) {
	if got := FromDateTime(2026, time.June, 30, 14, 0, 0).Format("hh A"); got != "02 PM" {
		t.Fatalf("Format(hh A) = %q, want 02 PM", got)
	}
}

func TestFormatWeekday(t *testing.T) {
	d := From(2026, time.June, 29)
	if got := d.Format("dddd"); got != "Monday" {
		t.Fatalf("Format(dddd) = %q, want Monday", got)
	}
	if got := d.Format("ddd"); got != "Mon" {
		t.Fatalf("Format(ddd) = %q, want Mon", got)
	}
}

func TestToDateString(t *testing.T) {
	if got := From(2026, time.June, 30).ToDateString(); got != "2026-06-30" {
		t.Fatalf("ToDateString() = %q, want 2026-06-30", got)
	}
}

func TestToTimeString(t *testing.T) {
	if got := FromDateTime(2026, time.June, 30, 10, 20, 30).ToTimeString(); got != "10:20:30" {
		t.Fatalf("ToTimeString() = %q, want 10:20:30", got)
	}
}

func TestToDateTimeString(t *testing.T) {
	if got := FromDateTime(2026, time.June, 30, 10, 20, 30).ToDateTimeString(); got != "2026-06-30 10:20:30" {
		t.Fatalf("ToDateTimeString() = %q, want 2026-06-30 10:20:30", got)
	}
}

func TestToFriendly(t *testing.T) {
	if got := From(2026, time.June, 29).ToFriendly(); got != "Monday, June 29 2026" {
		t.Fatalf("ToFriendly() = %q, want Monday, June 29 2026", got)
	}
}

func TestToRFC3339(t *testing.T) {
	if got := From(2026, time.June, 30).ToRFC3339(); got == "" {
		t.Fatal("ToRFC3339() must not be empty")
	}
}

func TestParse(t *testing.T) {
	cases := []string{
		"2026-06-30",
		"2026-06-30 15:04:05",
		"06/30/2026",
		"June 30, 2026",
		"Jun 30, 2026",
		"30 June 2026",
	}
	for _, s := range cases {
		if _, err := Parse(s); err != nil {
			t.Errorf("Parse(%q) unexpected error: %v", s, err)
		}
	}
}

func TestParseInvalid(t *testing.T) {
	if _, err := Parse("not a date"); err == nil {
		t.Fatal("Parse() must return error for unparseable input")
	}
}

func TestMustParse(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("MustParse() panicked on valid input: %v", r)
		}
	}()
	MustParse("2026-06-30")
}

func TestMustParsePanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("MustParse() must panic on invalid input")
		}
	}()
	MustParse("not a date")
}

func TestHumanize(t *testing.T) {
	cases := []struct {
		offset time.Duration
		want   string
	}{
		{0, "just now"},
		{-30 * time.Second, "30 seconds ago"},
		{-2 * time.Minute, "2 minutes ago"},
		{-5 * time.Hour, "5 hours ago"},
		{-3 * Day, "3 days ago"},
		{-2 * Week, "2 weeks ago"},
		{-2 * Month, "2 months ago"},
		{-2 * Year, "2 years ago"},
		{5 * time.Hour, "in 5 hours"},
	}
	for _, c := range cases {
		if got := New(time.Now().Add(c.offset)).Humanize(); got != c.want {
			t.Errorf("Humanize(%v) = %q, want %q", c.offset, got, c.want)
		}
	}
}

func TestFromNow(t *testing.T) {
	if got := New(time.Now().Add(-time.Minute)).FromNow(); got != "1 minute ago" {
		t.Fatalf("FromNow() = %q, want 1 minute ago", got)
	}
}

func TestDaysUntil(t *testing.T) {
	got := New(time.Now().Add(3 * Day)).DaysUntil()
	if got != 2 && got != 3 {
		t.Fatalf("DaysUntil() = %d, want 2 or 3", got)
	}
}

func TestDaysSince(t *testing.T) {
	got := New(time.Now().Add(-3 * Day)).DaysSince()
	if got != 2 && got != 3 {
		t.Fatalf("DaysSince() = %d, want 2 or 3", got)
	}
}

func TestDiff(t *testing.T) {
	if got := From(2026, time.June, 30).Diff(From(2026, time.June, 28)); got != "2 days" {
		t.Fatalf("Diff() = %q, want 2 days", got)
	}
}

func TestIsBusinessDay(t *testing.T) {
	if !From(2026, time.June, 29).IsBusinessDay() {
		t.Fatal("Monday must be a business day")
	}
	if From(2026, time.June, 27).IsBusinessDay() {
		t.Fatal("Saturday must not be a business day")
	}
}

func TestNextBusinessDay(t *testing.T) {
	if got := From(2026, time.June, 26).NextBusinessDay().Weekday(); got != time.Monday {
		t.Fatalf("NextBusinessDay() from Friday = %s, want Monday", got)
	}
}

func TestPrevBusinessDay(t *testing.T) {
	if got := From(2026, time.June, 29).PrevBusinessDay().Weekday(); got != time.Friday {
		t.Fatalf("PrevBusinessDay() from Monday = %s, want Friday", got)
	}
}

func TestAddBusinessDays(t *testing.T) {
	if got := From(2026, time.June, 26).AddBusinessDays(1).Weekday(); got != time.Monday {
		t.Fatalf("AddBusinessDays(1) from Friday = %s, want Monday", got)
	}
}

func TestAddBusinessDaysNegative(t *testing.T) {
	if got := From(2026, time.June, 29).AddBusinessDays(-1).Weekday(); got != time.Friday {
		t.Fatalf("AddBusinessDays(-1) from Monday = %s, want Friday", got)
	}
}

func TestBusinessDaysBetween(t *testing.T) {
	friday := From(2026, time.June, 26)
	monday := From(2026, time.June, 29)
	cases := []struct {
		a, b T
		want int
	}{
		{friday, monday, 1},
		{monday, friday, 1},
	}
	for _, c := range cases {
		if got := c.a.BusinessDaysBetween(c.b); got != c.want {
			t.Errorf("BusinessDaysBetween() = %d, want %d", got, c.want)
		}
	}
}

func TestUTC(t *testing.T) {
	if Now().UTC().Location() != time.UTC {
		t.Fatal("UTC() must return UTC location")
	}
}

func TestLocal(t *testing.T) {
	if Now().UTC().Local().Location() == time.UTC {
		t.Fatal("Local() must not return UTC location")
	}
}

func TestInLocation(t *testing.T) {
	ny, err := Now().InLocation("America/New_York")
	if err != nil {
		t.Fatalf("InLocation() error: %v", err)
	}
	if ny.Location().String() != "America/New_York" {
		t.Fatalf("InLocation() = %s, want America/New_York", ny.Location())
	}
}

func TestInLocationInvalid(t *testing.T) {
	if _, err := Now().InLocation("Not/Real"); err == nil {
		t.Fatal("InLocation() must return error for invalid timezone")
	}
}

func TestMustInLocation(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("MustInLocation() panicked on valid timezone: %v", r)
		}
	}()
	Now().MustInLocation("Europe/London")
}

func TestMustInLocationPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("MustInLocation() must panic on invalid timezone")
		}
	}()
	Now().MustInLocation("Not/Real")
}

func TestOffsetHours(t *testing.T) {
	if got := Now().UTC().OffsetHours(); got != 0 {
		t.Fatalf("OffsetHours() on UTC = %d, want 0", got)
	}
}
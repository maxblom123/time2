# time2

![CI](https://github.com/maxblom123/time2/actions/workflows/ci.yml/badge.svg)
![Go Version](https://img.shields.io/badge/go-1.26.4+-00ADD8?logo=go)
![Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)
![License](https://img.shields.io/badge/license-MIT-blue)

A production-grade companion to Go's standard `time` package. Built out of frustration with Go's reference-time formatting, missing duration constants, and the absence of any humanization layer — `time2` fills those gaps without replacing anything. Every `time.Time` method is still available; `time2` only adds.

```go
t := time2.Now()

t.Format("YYYY-MM-DD")           // "2026-06-30"  — not "2006-01-02"
t.Humanize()                     // "3 hours ago"
t.AddBusinessDays(5)             // skips Saturday and Sunday
t.StartOfWeek()                  // Monday 00:00:00
time2.Parse("June 30, 2026")     // no layout argument required
```

---

## Motivation

Go's `time` package is correct and well-designed. It is not ergonomic.

```go
// Standard library: format a date
t.Format("2006-01-02 15:04:05")

// time2: format a date
t.Format("YYYY-MM-DD HH:mm:ss")
```

The reference-time approach (`Mon Jan 2 15:04:05 MST 2006`) is clever in theory, each field maps to a unique value, but it is a source of constant bugs in practice. Developers copy-paste layouts from Stack Overflow, misremember whether month is `01` or `1`, and spend time reading documentation that should be unnecessary.

Beyond formatting, the standard library has no concept of a business day, no humanized duration output, no multi-layout parser, and no convenience constants for `Day`, `Week`, or `Month`. These are not exotic requirements, they appear in virtually every production application that does anything with time.

`time2` is an honest attempt to fix that with zero dependencies and 100% test coverage.

---

## Installation

```bash
go get github.com/maxblom123/time2
```

Requires Go 1.26.4 or later.

---

## API

### Constructors

```go
time2.Now()                                         // current local time
time2.New(t time.Time)                              // wrap an existing time.Time
time2.From(2026, time.June, 30)                     // date only, midnight local
time2.FromDateTime(2026, time.June, 30, 14, 0, 0)   // full datetime
time2.Parse("June 30, 2026")                        // automatic layout detection
time2.MustParse("2026-06-30")                       // panics on failure
```

All constructors return `time2.T`, which embeds `time.Time`, every standard library method works without any conversion.

---

### Formatting

`time2` uses a pattern-based layout system instead of Go's reference time. Patterns are case-sensitive and composable.

| Token  | Output            | Example  |
|--------|-------------------|----------|
| `YYYY` | 4-digit year      | `2026`   |
| `YY`   | 2-digit year      | `26`     |
| `MMMM` | Full month name   | `June`   |
| `MMM`  | Short month name  | `Jun`    |
| `MM`   | Zero-padded month | `06`     |
| `DD`   | Zero-padded day   | `30`     |
| `HH`   | 24-hour hour      | `14`     |
| `hh`   | 12-hour hour      | `02`     |
| `mm`   | Minutes           | `05`     |
| `ss`   | Seconds           | `09`     |
| `A`    | AM/PM             | `PM`     |
| `dddd` | Full weekday      | `Monday` |
| `ddd`  | Short weekday     | `Mon`    |

```go
t := time2.From(2026, time.June, 30)

t.Format("YYYY-MM-DD")            // "2026-06-30"
t.Format("dddd, MMMM DD YYYY")    // "Tuesday, June 30 2026"
t.Format("DD/MM/YYYY")            // "30/06/2026"
```

**Presets**

```go
t.ToDateString()      // "2026-06-30"
t.ToTimeString()      // "14:05:09"
t.ToDateTimeString()  // "2026-06-30 14:05:09"
t.ToFriendly()        // "Tuesday, June 30 2026"
t.ToRFC3339()         // "2026-06-30T14:05:09+02:00"
```

---

### Parsing

`Parse` tries eleven common layouts in order and returns the first match. No layout argument required.

```go
time2.Parse("2026-06-30")
time2.Parse("2026-06-30 14:05:09")
time2.Parse("06/30/2026")
time2.Parse("June 30, 2026")
time2.Parse("Jun 30, 2026")
time2.Parse("30 June 2026")
```

Returns `(T, error)`. Use `MustParse` in tests and package-level `var` declarations where a parse failure is a programming error, not a runtime condition.

---

### Predicates

```go
t.IsPast()        // true if t is before time.Now()
t.IsFuture()      // true if t is after time.Now()
t.IsToday()       // true if t falls on the current calendar day
t.IsWeekend()     // true if t is Saturday or Sunday
t.IsWeekday()     // true if t is Monday through Friday
t.IsBusinessDay() // alias for IsWeekday()
```

---

### Arithmetic

```go
t.AddDays(n)    // add n calendar days (negative to subtract)
t.AddWeeks(n)   // add n weeks
t.AddMonths(n)  // add n months (uses time.AddDate — handles month-end correctly)
t.AddYears(n)   // add n years
```

**Business day arithmetic** skips Saturday and Sunday in both directions:

```go
friday := time2.From(2026, time.June, 26)

friday.AddBusinessDays(1)   // Monday June 29
friday.AddBusinessDays(-1)  // Thursday June 25
friday.NextBusinessDay()    // Monday June 29
friday.PrevBusinessDay()    // Thursday June 25

monday := time2.From(2026, time.June, 29)
friday.BusinessDaysBetween(monday)  // 1
```

---

### Period boundaries

```go
t.StartOfDay()    // 00:00:00.000000000
t.EndOfDay()      // 23:59:59.999999999
t.StartOfWeek()   // Monday 00:00:00 of the containing week
t.StartOfMonth()  // 1st of the month at 00:00:00
t.EndOfMonth()    // last nanosecond of the month
t.StartOfYear()   // January 1st at 00:00:00
```

---

### Humanization

```go
time2.Now().Humanize()                                         // "just now"
time2.New(time.Now().Add(-90*time.Second)).Humanize()          // "1 minute ago"
time2.New(time.Now().Add(-2*time.Hour)).Humanize()             // "2 hours ago"
time2.New(time.Now().Add(-3*time2.Day)).Humanize()             // "3 days ago"
time2.New(time.Now().Add(24*time.Hour)).Humanize()             // "in 1 day"

t.FromNow()    // alias for Humanize()
t.DaysSince()  // integer count of days elapsed since t
t.DaysUntil()  // integer count of days until t (negative if past)
t.Diff(u)      // humanized duration between t and u
```

`HumanizeDuration` is also exported for use with arbitrary `time.Duration` values:

```go
time2.HumanizeDuration(90 * time.Minute)  // "1 hour 30 minutes"
time2.HumanizeDuration(48 * time.Hour)    // "2 days"
```

---

### Duration constants

The standard library defines `time.Hour` but stops there. `time2` adds the rest:

```go
time2.Second  // time.Second
time2.Minute  // time.Minute
time2.Hour    // time.Hour
time2.Day     // 24 * time.Hour
time2.Week    // 7 * time2.Day
time2.Month   // 30 * time2.Day  (approximate)
time2.Year    // 365 * time2.Day (approximate)
```

---

### Timezones

```go
t.UTC()                            // convert to UTC
t.Local()                          // convert to local timezone
t.InLocation("America/New_York")   // returns (T, error)
t.MustInLocation("Europe/London")  // panics on invalid timezone
t.OffsetHours()                    // UTC offset as integer, e.g. +2 or -5
```

---

## Design decisions

**`T` embeds `time.Time` rather than wrapping it.**
This means `time2.T` satisfies any interface that accepts `time.Time` methods and requires no conversion when passing values to standard library functions. The tradeoff is that `Format` shadows `time.Time.Format` — callers who need Go's reference-time formatting can access it via `t.Time.Format(layout)`.

**No external dependencies.**
The module graph is exactly one node. There is nothing to audit, nothing to update, nothing that can introduce a vulnerability through a transitive dependency.

**`MustParse` is intentional.**
Panic-on-failure constructors are appropriate when the input is a compile-time constant and a failure is a programmer error. `MustParse` follows the same convention as `regexp.MustCompile` and `template.Must`.

**Business day logic does not account for public holidays.**
Holiday calendars are locale-specific and change over time. Including one would either require an external data source or hard-code assumptions that break for international users. The current implementation is explicit about what it does: it skips weekends. Callers who need holiday awareness can compose `AddBusinessDays` with their own exclusion list.

**Approximate month and year constants.**
`time2.Month` and `time2.Year` are duration constants, not calendar concepts. They are useful for rough comparisons and humanization thresholds, not for precise date arithmetic. Precise month and year arithmetic should use `AddMonths` and `AddYears`, which delegate to `time.AddDate`.

---

## Testing

```bash
go test ./...
go test -race ./...
go test "-coverprofile=coverage.out" .
go tool cover "-func=coverage.out"
```

Coverage is 100% of statements. The race detector passes cleanly.

---

## Security and static analysis

```bash
govulncheck ./...   # no vulnerabilities
gosec ./...         # no findings
staticcheck ./...   # no findings
go vet ./...        # no findings
```

---

## License

MIT

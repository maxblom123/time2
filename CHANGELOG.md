# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2026-06-30

### Added
- `T` type embedding `time.Time` — full standard library compatibility
- Constructors: `Now`, `New`, `From`, `FromDateTime`
- Sane formatting with `YYYY-MM-DD` style patterns instead of Go reference time
- Format presets: `ToDateString`, `ToTimeString`, `ToDateTimeString`, `ToFriendly`, `ToRFC3339`
- Smart multi-layout parser: `Parse`, `MustParse`
- Predicates: `IsPast`, `IsFuture`, `IsToday`, `IsWeekend`, `IsWeekday`, `IsBusinessDay`
- Calendar arithmetic: `AddDays`, `AddWeeks`, `AddMonths`, `AddYears`
- Business day arithmetic: `AddBusinessDays`, `NextBusinessDay`, `PrevBusinessDay`, `BusinessDaysBetween`
- Period boundaries: `StartOfDay`, `EndOfDay`, `StartOfWeek`, `StartOfMonth`, `EndOfMonth`, `StartOfYear`
- Humanization: `Humanize`, `FromNow`, `DaysUntil`, `DaysSince`, `Diff`, `HumanizeDuration`
- Duration constants: `Second`, `Minute`, `Hour`, `Day`, `Week`, `Month`, `Year`
- Timezone helpers: `UTC`, `Local`, `InLocation`, `MustInLocation`, `OffsetHours`
- 100% test coverage
- Zero external dependencies
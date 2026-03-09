# XTime Package - Calendar And Schedule Notes

**Package:** `github.com/lazygophers/utils/xtime`

## OVERVIEW
- `xtime` is not just time formatting; it contains calendar-domain logic.
- Main areas: custom time constants/types, lunar calendar conversion, solar-term helpers, and work-schedule subpackages.
- The package has meaningful sub-boundaries: `xtime007/`, `xtime955/`, `xtime996/`.

## WHERE TO LOOK
- `xtime.go`: shared duration/time constants.
- `now.go`, `time.go`: top-level time wrappers/helpers.
- `calendar.go`: `Calendar` and `NowCalendar()` entrypoint.
- `lunar.go`: `Lunar`, `WithLunarTime`, `WithLunar`, solar-to-lunar conversion logic.
- `solarterm.go`: solar term calculations.
- `xtime007/`, `xtime955/`, `xtime996/`: schedule-specific business rules.

## LOCAL RULES
- Treat lunar/calendar code as domain logic, not simple date arithmetic.
- Preserve the separation between generic time helpers and schedule-specific subpackages.
- If changing exported calendar behavior, verify both alias-style output (`Animal`, `YearAlias`, `MonthAlias`, `DayAlias`) and raw date accessors.
- If changing conversion logic, review helper functions such as `daysOfLunarYear`, `leapMonth`, `leapDays`, and `lunarDays` together.

## GOTCHAS
- Existing `llms.txt` in this package is too generic in parts; prefer source files.
- `xtime007/955/996` are not cosmetic folders; they encode different schedule assumptions.
- Calendar output includes Chinese-domain concepts like zodiac and lunar aliases, so localization-sensitive changes need extra care.

## TESTING
```bash
go test ./xtime/...
make test
```

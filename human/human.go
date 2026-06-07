// Package human formats raw values (byte sizes, speeds, durations, relative
// times) into human-readable strings with locale-aware units.
//
// Locale resolution follows github.com/lazygophers/utils/language:
// goroutine-local override → global default → "en". Per-package toggles
// (precision, compact, clock format) live as package-level state and have
// dedicated Set* setters.
package human

import (
	"math"
	"strconv"
	"strings"

	"github.com/lazygophers/utils/language"
)

// Package-level rendering state.
var (
	defaultPrecision   = 1
	defaultCompact     = false
	defaultClockFormat = false
)

// SetDefaultPrecision sets the number of fractional digits kept on
// fractional values.
func SetDefaultPrecision(precision int) { defaultPrecision = precision }

// SetCompact toggles the compact rendering mode (no space between value and unit).
func SetCompact(compact bool) { defaultCompact = compact }

// SetClockFormat toggles the clock-style rendering for durations (HH:MM:SS).
func SetClockFormat(clock bool) { defaultClockFormat = clock }

// currentLocaleName returns the goroutine-effective locale base code
// (e.g. "zh", "en"). Falls back to the global default, then "en".
func currentLocaleName() string {
	t := language.Get()
	if t == nil {
		return defaultLocaleName()
	}
	base, _ := t.Tag().Base()
	return base.String()
}

// defaultLocaleName returns the global default locale base code.
func defaultLocaleName() string {
	t := language.Default()
	if t == nil {
		return "en"
	}
	base, _ := t.Tag().Base()
	return base.String()
}

// formatWithUnit composes a value and unit string subject to compact mode and
// precision settings.
func formatWithUnit(value float64, unitIndex int, unitType string) string {
	locale, _ := GetLocaleConfig(currentLocaleName())

	var units []string
	switch unitType {
	case "byte":
		units = locale.ByteUnits
	case "speed":
		units = locale.SpeedUnits
	case "bitspeed":
		units = locale.BitSpeedUnits
	default:
		return "-"
	}

	if unitIndex >= len(units) {
		unitIndex = len(units) - 1
	}

	unit := units[unitIndex]

	var formattedValue string
	if value == math.Trunc(value) {
		formattedValue = strconv.FormatFloat(value, 'f', 0, 64)
	} else {
		precision := defaultPrecision
		if precision < 0 {
			precision = 1
		}
		formattedValue = formatFloat(value, precision)
	}

	if defaultCompact {
		return formattedValue + unit
	}

	return formattedValue + " " + unit
}

// formatFloat formats f with trailing zero / dot trimming.
func formatFloat(f float64, precision int) string {
	if precision < 0 {
		precision = 0
	}

	str := strconv.FormatFloat(f, 'f', precision, 64)

	if strings.Contains(str, ".") {
		str = strings.TrimRight(str, "0")
		str = strings.TrimRight(str, ".")
	}

	return str
}

// abs returns the absolute value of an int64.
func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

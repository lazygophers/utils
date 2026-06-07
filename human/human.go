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

	xlanguage "golang.org/x/text/language"

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

// currentTag returns the goroutine-effective language tag. Resolution order:
// goroutine-local override → global default → English.
func currentTag() xlanguage.Tag {
	t := language.Get()
	if t == nil {
		t = language.Default()
	}
	if t == nil {
		return xlanguage.English
	}
	return t.Tag()
}

// formatScaled scales v down by base until it fits a unit slot, then composes
// the value + unit. Shared by ByteSize / Speed / BitSpeed.
func formatScaled(v int64, base int64, units []string) string {
	if v == 0 || len(units) == 0 {
		return formatValueWithUnit(0, units, 0)
	}

	// Iterative integer division avoids math.Log + math.Pow per call.
	absV := v
	if absV < 0 {
		absV = -absV
	}
	exp := 0
	divisor := int64(1)
	for absV >= base && exp < len(units)-1 {
		absV /= base
		divisor *= base
		exp++
	}

	value := float64(v) / float64(divisor)
	return formatValueWithUnit(value, units, exp)
}

// formatValueWithUnit composes a numeric value and the unit at unitIndex
// inside units, honoring compact mode and the current precision setting.
func formatValueWithUnit(value float64, units []string, unitIndex int) string {
	if len(units) == 0 {
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

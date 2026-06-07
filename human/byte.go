package human

import "math"

// ByteSize formats a raw byte count using binary (1024) scaling and the
// current locale's byte units.
func ByteSize(bytes int64) string { return formatByteSize(bytes) }

func formatByteSize(bytes int64) string {
	locale, _ := GetLocaleConfig(currentLocaleName())
	units := locale.ByteUnits

	if bytes == 0 {
		return formatByteUnit(0, 0, units)
	}

	absBytes := abs(bytes)
	const unit = 1024

	if absBytes < unit {
		return formatByteUnit(float64(bytes), 0, units)
	}

	exp := int(math.Floor(math.Log(float64(absBytes)) / math.Log(unit)))
	if exp >= len(units) {
		exp = len(units) - 1
	}

	value := float64(bytes) / math.Pow(unit, float64(exp))
	return formatByteUnit(value, exp, units)
}

// formatByteUnit renders a byte value with the byte unit at unitIndex.
func formatByteUnit(value float64, unitIndex int, units []string) string {
	return formatValueWithUnit(value, units, unitIndex)
}

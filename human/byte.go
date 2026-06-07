package human

import "math"

// ByteSize formats a raw byte count using binary (1024) scaling and the
// current locale's byte units.
func ByteSize(bytes int64) string { return formatByteSize(bytes) }

func formatByteSize(bytes int64) string {
	if bytes == 0 {
		return formatWithUnit(0, 0, "byte")
	}

	absBytes := abs(bytes)
	const unit = 1024

	locale, _ := GetLocaleConfig(currentLocaleName())
	units := locale.ByteUnits

	if absBytes < unit {
		return formatWithUnit(float64(bytes), 0, "byte")
	}

	exp := int(math.Floor(math.Log(float64(absBytes)) / math.Log(unit)))
	if exp >= len(units) {
		exp = len(units) - 1
	}

	value := float64(bytes) / math.Pow(unit, float64(exp))
	return formatWithUnit(value, exp, "byte")
}

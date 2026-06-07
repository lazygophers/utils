package human

import "math"

// Speed formats a bytes-per-second rate using binary (1024) scaling and
// the current locale's speed units.
func Speed(bytesPerSecond int64) string { return formatSpeed(bytesPerSecond) }

// BitSpeed formats a bits-per-second rate using decimal (1000) scaling and
// the current locale's bit-speed units. Decimal scaling matches network
// conventions.
func BitSpeed(bitsPerSecond int64) string { return formatBitSpeed(bitsPerSecond) }

func formatSpeed(bytesPerSecond int64) string {
	locale, _ := GetLocaleConfig(currentLocaleName())
	units := locale.SpeedUnits

	if bytesPerSecond == 0 {
		return formatSpeedUnit(0, 0, units)
	}

	absBytes := abs(bytesPerSecond)
	const unit = 1024

	if absBytes < unit {
		return formatSpeedUnit(float64(bytesPerSecond), 0, units)
	}

	exp := int(math.Floor(math.Log(float64(absBytes)) / math.Log(unit)))
	if exp >= len(units) {
		exp = len(units) - 1
	}

	value := float64(bytesPerSecond) / math.Pow(unit, float64(exp))
	return formatSpeedUnit(value, exp, units)
}

func formatBitSpeed(bitsPerSecond int64) string {
	locale, _ := GetLocaleConfig(currentLocaleName())
	units := locale.BitSpeedUnits

	if bitsPerSecond == 0 {
		return formatBitSpeedUnit(0, 0, units)
	}

	absBits := abs(bitsPerSecond)
	const unit = 1000

	if absBits < unit {
		return formatBitSpeedUnit(float64(bitsPerSecond), 0, units)
	}

	exp := int(math.Floor(math.Log(float64(absBits)) / math.Log(unit)))
	if exp >= len(units) {
		exp = len(units) - 1
	}

	value := float64(bitsPerSecond) / math.Pow(unit, float64(exp))
	return formatBitSpeedUnit(value, exp, units)
}

// formatSpeedUnit renders a byte-per-second value with the speed unit.
func formatSpeedUnit(value float64, unitIndex int, units []string) string {
	return formatValueWithUnit(value, units, unitIndex)
}

// formatBitSpeedUnit renders a bit-per-second value with the bit-speed unit.
func formatBitSpeedUnit(value float64, unitIndex int, units []string) string {
	return formatValueWithUnit(value, units, unitIndex)
}

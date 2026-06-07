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
	if bytesPerSecond == 0 {
		return formatWithUnit(0, 0, "speed")
	}

	absBytes := abs(bytesPerSecond)
	const unit = 1024

	locale, _ := GetLocaleConfig(currentLocaleName())
	units := locale.SpeedUnits

	if absBytes < unit {
		return formatWithUnit(float64(bytesPerSecond), 0, "speed")
	}

	exp := int(math.Floor(math.Log(float64(absBytes)) / math.Log(unit)))
	if exp >= len(units) {
		exp = len(units) - 1
	}

	value := float64(bytesPerSecond) / math.Pow(unit, float64(exp))
	return formatWithUnit(value, exp, "speed")
}

func formatBitSpeed(bitsPerSecond int64) string {
	if bitsPerSecond == 0 {
		return formatWithUnit(0, 0, "bitspeed")
	}

	absBits := abs(bitsPerSecond)
	const unit = 1000

	locale, _ := GetLocaleConfig(currentLocaleName())
	units := locale.BitSpeedUnits

	if absBits < unit {
		return formatWithUnit(float64(bitsPerSecond), 0, "bitspeed")
	}

	exp := int(math.Floor(math.Log(float64(absBits)) / math.Log(unit)))
	if exp >= len(units) {
		exp = len(units) - 1
	}

	value := float64(bitsPerSecond) / math.Pow(unit, float64(exp))
	return formatWithUnit(value, exp, "bitspeed")
}

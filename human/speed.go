package human

// Speed formats a bytes-per-second rate using binary (1024) scaling and
// the current locale's speed units.
func Speed(bytesPerSecond int64) string { return formatSpeed(bytesPerSecond) }

// BitSpeed formats a bits-per-second rate using decimal (1000) scaling and
// the current locale's bit-speed units. Decimal scaling matches network
// conventions.
func BitSpeed(bitsPerSecond int64) string { return formatBitSpeed(bitsPerSecond) }

func formatSpeed(bytesPerSecond int64) string {
	locale, _ := GetLocaleConfig(currentTag())
	return formatScaled(bytesPerSecond, 1024, locale.SpeedUnits)
}

func formatBitSpeed(bitsPerSecond int64) string {
	locale, _ := GetLocaleConfig(currentTag())
	return formatScaled(bitsPerSecond, 1000, locale.BitSpeedUnits)
}

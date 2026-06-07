package human

// ByteSize formats a raw byte count using binary (1024) scaling and the
// current locale's byte units.
func ByteSize(bytes int64) string { return formatByteSize(bytes) }

func formatByteSize(bytes int64) string {
	locale, _ := GetLocaleConfig(currentLocaleName())
	return formatScaled(bytes, 1024, locale.ByteUnits)
}

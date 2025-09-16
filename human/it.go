package human

// Bit units (base unit = 1 bit)
const (
	Bit  = 1
	Kbit = 1000 * Bit // Network speeds typically use decimal
	Mbit = 1000 * Kbit
	Gbit = 1000 * Mbit
	Tbit = 1000 * Gbit
	Pbit = 1000 * Tbit
	Ebit = 1000 * Pbit
	Zbit = 1000 * Ebit
)

// Byte units (base unit = 8 bits = 1 byte)
const (
	Byte = 8    // 1 byte = 8 bits
	KB   = 1024 // 1 KB = 1024 bytes (binary)
	MB   = 1024 * KB
	GB   = 1024 * MB
	TB   = 1024 * GB
	PB   = 1024 * TB
	EB   = 1024 * PB
	ZB   = 1024 * EB
)

// Alternative decimal byte units (for compatibility)
const (
	KBDecimal = 1000 * Byte // 1 KB = 1000 bytes (decimal)
	MBDecimal = 1000 * KBDecimal
	GBDecimal = 1000 * MBDecimal
	TBDecimal = 1000 * GBDecimal
	PBDecimal = 1000 * TBDecimal
	EBDecimal = 1000 * PBDecimal
	ZBDecimal = 1000 * EBDecimal
)

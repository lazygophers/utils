package human

// Bit units (base unit = 1 bit)
const (
	Bit = 1
	Kb  = 1000 * Bit // Network speeds typically use decimal
	Mb  = 1000 * Kb
	Gb  = 1000 * Mb
	Tb  = 1000 * Gb
	Pb  = 1000 * Tb
	Eb  = 1000 * Pb
	Zb  = 1000 * Eb
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

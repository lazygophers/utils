package candy

// ValueType represents the type of a value for type checking
type ValueType int

const (
	// ValueUnknown represents an unknown or unsupported type
	ValueUnknown ValueType = iota
	// ValueNumber represents numeric types (int, float, etc.)
	ValueNumber
	// ValueString represents string and byte slice types
	ValueString
	// ValueBool represents boolean type
	ValueBool
)

// CheckValueType determines the general category of a value's type
func CheckValueType(val interface{}) ValueType {
	switch val.(type) {
	case bool:
		return ValueBool
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return ValueNumber
	case string, []byte:
		return ValueString
	default:
		return ValueUnknown
	}
}
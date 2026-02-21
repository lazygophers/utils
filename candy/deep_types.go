package candy

// Comparable 定义可比较的类型约束
type Comparable interface {
	comparable
}

// Copyable 定义可复制的基本类型约束
type Copyable interface {
	~bool | ~string |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~complex64 | ~complex128
}

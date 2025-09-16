package candy

func Each[T any](values []T, fn func(value T)) {
	for _, value := range values {
		fn(value)
	}
}

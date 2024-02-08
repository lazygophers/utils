package json

func MustMarshal(v any) []byte {
	buf, err := Marshal(v)
	if err != nil {
		panic(err)
	}
	return buf
}

func MustMarshalString(v any) string {
	buf, err := MarshalString(v)
	if err != nil {
		panic(err)
	}
	return buf
}

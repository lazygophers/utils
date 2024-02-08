package json

import "os"

func UnmarshalFromFile(filename string, v any) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return NewDecoder(file).Decode(v)
}

func MustUnmarshalFromFile(filename string, v any) {
	err := UnmarshalFromFile(filename, v)
	if err != nil {
		panic(err)
	}
}

func MarshalToFile(filename string, v any) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return NewEncoder(file).Encode(v)
}

func MustMarshalToFile(filename string, v any) {
	err := MarshalToFile(filename, v)
	if err != nil {
		panic(err)
	}
}

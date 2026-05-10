package cryptox

import (
	"github.com/google/uuid"
)

func UUID() string {
	s := uuid.New().String()
	var result [32]byte
	copy(result[0:8], s[0:8])
	copy(result[8:12], s[9:13])
	copy(result[12:16], s[14:18])
	copy(result[16:20], s[19:23])
	copy(result[20:32], s[24:36])
	return string(result[:])
}

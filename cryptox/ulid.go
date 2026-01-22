package cryptox

import (
	"github.com/lazygophers/log"
	"github.com/lazygophers/utils"
	"github.com/oklog/ulid/v2"
)

func ULID() string {
	return ulid.Make().String()
}

func ULIDWithTimestamp() (string, int64) {
	id := ulid.Make()
	return id.String(), int64(id.Time())
}

// GetULIDTimestamp 从 ULID 中提取时间戳（毫秒级）
func GetULIDTimestamp(id string) (int64, error) {
	parsedId, err := ulid.Parse(id)
	if err != nil {
		log.Errorf("err:%v", err)
		return 0, err
	}

	// ULID 时间戳以毫秒为单位
	return int64(parsedId.Time()), nil
}

// MustGetULIDTimestamp 从 ULID 中提取时间戳（毫秒级）
func MustGetULIDTimestamp(id string) int64 {
	return utils.Must(GetULIDTimestamp(id))
}

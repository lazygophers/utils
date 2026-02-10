package app

import (
	"os"
	"strings"
)

var (
	Commit      string
	ShortCommit string
	Branch      string
	Tag         string

	BuildDate string

	GoVersion string

	GoOS    string
	Goarch  string
	Goarm   string
	Goamd64 string
	Gomips  string

	Description string
)

// 通过环境变量设置日志级别
func setPackageTypeFromEnv() {
	switch strings.ToLower(os.Getenv("APP_ENV")) {
	case "dev", "development":
		Env = Debug
	case "test", "canary":
		Env = Test
	case "prod", "release", "production":
		Env = Release
	case "alpha":
		Env = Alpha
	case "beta":
		Env = Beta
	}
}

func init() {
	setPackageTypeFromEnv()
}

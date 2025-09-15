//go:build lang_ko || lang_all

package fake

import "embed"

//go:embed data/ko
var dataKO embed.FS

func init() {
	RegisterLanguageFS("ko", dataKO)
}
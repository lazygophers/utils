//go:build lang_de || lang_all

package fake

import "embed"

//go:embed data/de
var dataDE embed.FS

func init() {
	RegisterLanguageFS("de", dataDE)
}

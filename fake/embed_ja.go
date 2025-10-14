//go:build lang_ja || lang_all

package fake

import "embed"

//go:embed data/ja
var dataJA embed.FS

func init() {
	RegisterLanguageFS("ja", dataJA)
}

//go:build lang_en || lang_all

package fake

import "embed"

//go:embed data/en
var dataEN embed.FS

func init() {
	RegisterLanguageFS("en", dataEN)
}

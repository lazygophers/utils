//go:build lang_it || lang_all

package fake

import "embed"

//go:embed data/it
var dataIT embed.FS

func init() {
	RegisterLanguageFS("it", dataIT)
}

//go:build lang_es || lang_all

package fake

import "embed"

//go:embed data/es
var dataES embed.FS

func init() {
	RegisterLanguageFS("es", dataES)
}

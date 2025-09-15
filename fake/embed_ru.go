//go:build lang_ru || lang_all

package fake

import "embed"

//go:embed data/ru
var dataRU embed.FS

func init() {
	RegisterLanguageFS("ru", dataRU)
}

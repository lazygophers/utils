//go:build fake_ru

package fake

import "embed"

//go:embed data/ru
var dataRU embed.FS

func init() {
	RegisterLanguageFS("ru", dataRU)
}

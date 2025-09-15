//go:build lang_fr || lang_all

package fake

import "embed"

//go:embed data/fr
var dataFR embed.FS

func init() {
	RegisterLanguageFS("fr", dataFR)
}

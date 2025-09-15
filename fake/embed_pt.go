//go:build lang_pt || lang_all

package fake

import "embed"

//go:embed data/pt
var dataPT embed.FS

func init() {
	RegisterLanguageFS("pt", dataPT)
}

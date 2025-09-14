//go:build fake_pt

package fake

import "embed"

//go:embed data/pt
var dataPT embed.FS

func init() {
	RegisterLanguageFS("pt", dataPT)
}
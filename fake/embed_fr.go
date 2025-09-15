//go:build fake_fr

package fake

import "embed"

//go:embed data/fr
var dataFR embed.FS

func init() {
	RegisterLanguageFS("fr", dataFR)
}

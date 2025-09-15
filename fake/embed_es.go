//go:build fake_es

package fake

import "embed"

//go:embed data/es
var dataES embed.FS

func init() {
	RegisterLanguageFS("es", dataES)
}

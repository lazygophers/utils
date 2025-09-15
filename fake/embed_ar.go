//go:build lang_ar || lang_all

package fake

import "embed"

//go:embed data/ar
var dataAR embed.FS

func init() {
	RegisterLanguageFS("ar", dataAR)
}
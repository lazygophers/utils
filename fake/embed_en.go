//go:build fake_en

package fake

import "embed"

//go:embed data/en
var dataEN embed.FS

func init() {
	RegisterLanguageFS("en", dataEN)
}

//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurksAndCaicosIslands.RegisterName(xlanguage.French, "Îles Turques-et-Caïques")
	dataTurksAndCaicosIslands.RegisterOfficialName(xlanguage.French, "Îles Turques-et-Caïques")
	dataTurksAndCaicosIslands.RegisterCapital(xlanguage.French, "Cockburn Town")
}

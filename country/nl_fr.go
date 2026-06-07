//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNetherlands.RegisterName(xlanguage.French, "Pays-Bas")
	dataNetherlands.RegisterOfficialName(xlanguage.French, "Royaume des Pays-Bas")
	dataNetherlands.RegisterCapital(xlanguage.French, "Amsterdam")
}

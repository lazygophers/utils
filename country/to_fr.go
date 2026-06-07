//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTonga.RegisterName(xlanguage.French, "Tonga")
	dataTonga.RegisterOfficialName(xlanguage.French, "Royaume des Tonga")
	dataTonga.RegisterCapital(xlanguage.French, "Nukuʻalofa")
}

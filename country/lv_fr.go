//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLatvia.RegisterName(xlanguage.French, "Lettonie")
	dataLatvia.RegisterOfficialName(xlanguage.French, "République de Lettonie")
	dataLatvia.RegisterCapital(xlanguage.French, "Riga")
}

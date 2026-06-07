//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAngola.RegisterName(xlanguage.French, "Angola")
	dataAngola.RegisterOfficialName(xlanguage.French, "République d'Angola")
	dataAngola.RegisterCapital(xlanguage.French, "Luanda")
}

//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGermany.RegisterName(xlanguage.French, "Allemagne")
	dataGermany.RegisterOfficialName(xlanguage.French, "République fédérale d'Allemagne")
	dataGermany.RegisterCapital(xlanguage.French, "Berlin")
}

//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMaldives.RegisterName(xlanguage.French, "Maldives")
	dataMaldives.RegisterOfficialName(xlanguage.French, "République des Maldives")
	dataMaldives.RegisterCapital(xlanguage.French, "Malé")
}

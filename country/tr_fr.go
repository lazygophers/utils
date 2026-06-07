//go:build (lang_fr || lang_all) && (country_all || country_asia || country_tr || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkey.RegisterName(xlanguage.French, "Turquie")
	dataTurkey.RegisterOfficialName(xlanguage.French, "République de Turquie")
	dataTurkey.RegisterCapital(xlanguage.French, "Ankara")
}

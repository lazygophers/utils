//go:build (lang_es || lang_all) && (country_all || country_asia || country_tr || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkey.RegisterName(xlanguage.Spanish, "Turquía")
	dataTurkey.RegisterOfficialName(xlanguage.Spanish, "República de Turquía")
	dataTurkey.RegisterCapital(xlanguage.Spanish, "Ankara")
}

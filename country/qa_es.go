//go:build (lang_es || lang_all) && (country_all || country_asia || country_qa || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataQatar.RegisterName(xlanguage.Spanish, "Catar")
	dataQatar.RegisterOfficialName(xlanguage.Spanish, "Estado de Catar")
	dataQatar.RegisterCapital(xlanguage.Spanish, "Doha")
}

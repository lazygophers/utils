//go:build (lang_es || lang_all) && (country_africa || country_all || country_gh || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGhana.RegisterName(xlanguage.Spanish, "Ghana")
	dataGhana.RegisterOfficialName(xlanguage.Spanish, "República de Ghana")
	dataGhana.RegisterCapital(xlanguage.Spanish, "Acra")
}

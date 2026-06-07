//go:build (lang_es || lang_all) && (country_africa || country_all || country_tg || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTogo.RegisterName(xlanguage.Spanish, "Togo")
	dataTogo.RegisterOfficialName(xlanguage.Spanish, "República Togolesa")
	dataTogo.RegisterCapital(xlanguage.Spanish, "Lomé")
}

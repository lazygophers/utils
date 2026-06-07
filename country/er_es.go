//go:build (lang_es || lang_all) && (country_africa || country_all || country_eastern_africa || country_er)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEritrea.RegisterName(xlanguage.Spanish, "Eritrea")
	dataEritrea.RegisterOfficialName(xlanguage.Spanish, "Estado de Eritrea")
	dataEritrea.RegisterCapital(xlanguage.Spanish, "Asmara")
}

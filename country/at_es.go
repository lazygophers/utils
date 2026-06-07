//go:build (lang_es || lang_all) && (country_all || country_at || country_europe || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustria.RegisterName(xlanguage.Spanish, "Austria")
	dataAustria.RegisterOfficialName(xlanguage.Spanish, "República de Austria")
	dataAustria.RegisterCapital(xlanguage.Spanish, "Viena")
}

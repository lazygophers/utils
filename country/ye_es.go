//go:build (lang_es || lang_all) && (country_all || country_asia || country_western_asia || country_ye)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataYemen.RegisterName(xlanguage.Spanish, "Yemen")
	dataYemen.RegisterOfficialName(xlanguage.Spanish, "República de Yemen")
	dataYemen.RegisterCapital(xlanguage.Spanish, "Saná")
}

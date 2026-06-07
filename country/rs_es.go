//go:build (lang_es || lang_all) && (country_all || country_europe || country_rs || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSerbia.RegisterName(xlanguage.Spanish, "Serbia")
	dataSerbia.RegisterOfficialName(xlanguage.Spanish, "República de Serbia")
	dataSerbia.RegisterCapital(xlanguage.Spanish, "Belgrado")
}

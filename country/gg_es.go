//go:build (lang_es || lang_all) && (country_all || country_europe || country_gg || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuernsey.RegisterName(xlanguage.Spanish, "Guernsey")
	dataGuernsey.RegisterOfficialName(xlanguage.Spanish, "Bailía de Guernsey")
	dataGuernsey.RegisterCapital(xlanguage.Spanish, "Saint Peter Port")
}

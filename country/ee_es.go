//go:build (lang_es || lang_all) && (country_all || country_ee || country_europe || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEstonia.RegisterName(xlanguage.Spanish, "Estonia")
	dataEstonia.RegisterOfficialName(xlanguage.Spanish, "República de Estonia")
	dataEstonia.RegisterCapital(xlanguage.Spanish, "Tallin")
}

//go:build (lang_es || lang_all) && (country_all || country_europe || country_ie || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIreland.RegisterName(xlanguage.Spanish, "Irlanda")
	dataIreland.RegisterOfficialName(xlanguage.Spanish, "Irlanda")
	dataIreland.RegisterCapital(xlanguage.Spanish, "Dublín")
}

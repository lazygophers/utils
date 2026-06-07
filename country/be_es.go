//go:build (lang_es || lang_all) && (country_all || country_be || country_europe || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelgium.RegisterName(xlanguage.Spanish, "Bélgica")
	dataBelgium.RegisterOfficialName(xlanguage.Spanish, "Reino de Bélgica")
	dataBelgium.RegisterCapital(xlanguage.Spanish, "Bruselas")
}

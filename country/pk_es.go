//go:build (lang_es || lang_all) && (country_all || country_asia || country_pk || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPakistan.RegisterName(xlanguage.Spanish, "Pakistán")
	dataPakistan.RegisterOfficialName(xlanguage.Spanish, "República Islámica de Pakistán")
	dataPakistan.RegisterCapital(xlanguage.Spanish, "Islamabad")
}

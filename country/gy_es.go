//go:build (lang_es || lang_all) && (country_all || country_americas || country_gy || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuyana.RegisterName(xlanguage.Spanish, "Guyana")
	dataGuyana.RegisterOfficialName(xlanguage.Spanish, "República Cooperativa de Guyana")
	dataGuyana.RegisterCapital(xlanguage.Spanish, "Georgetown")
}

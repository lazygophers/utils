//go:build (lang_es || lang_all) && (country_all || country_europe || country_sm || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSanMarino.RegisterName(xlanguage.Spanish, "San Marino")
	dataSanMarino.RegisterOfficialName(xlanguage.Spanish, "República de San Marino")
	dataSanMarino.RegisterCapital(xlanguage.Spanish, "San Marino")
}

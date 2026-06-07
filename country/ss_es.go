//go:build (lang_es || lang_all) && (country_africa || country_all || country_eastern_africa || country_ss)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthSudan.RegisterName(xlanguage.Spanish, "Sudán del Sur")
	dataSouthSudan.RegisterOfficialName(xlanguage.Spanish, "República de Sudán del Sur")
	dataSouthSudan.RegisterCapital(xlanguage.Spanish, "Yuba")
}

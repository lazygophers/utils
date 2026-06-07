//go:build (lang_es || lang_all) && (country_africa || country_all || country_northern_africa || country_sd)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSudan.RegisterName(xlanguage.Spanish, "Sudán")
	dataSudan.RegisterOfficialName(xlanguage.Spanish, "República de Sudán")
	dataSudan.RegisterCapital(xlanguage.Spanish, "Jartum")
}

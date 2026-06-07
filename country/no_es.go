//go:build (lang_es || lang_all) && (country_all || country_europe || country_no || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorway.RegisterName(xlanguage.Spanish, "Noruega")
	dataNorway.RegisterOfficialName(xlanguage.Spanish, "Reino de Noruega")
	dataNorway.RegisterCapital(xlanguage.Spanish, "Oslo")
}

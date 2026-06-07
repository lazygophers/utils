//go:build (lang_es || lang_all) && (country_all || country_europe || country_li || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLiechtenstein.RegisterName(xlanguage.Spanish, "Liechtenstein")
	dataLiechtenstein.RegisterOfficialName(xlanguage.Spanish, "Principado de Liechtenstein")
	dataLiechtenstein.RegisterCapital(xlanguage.Spanish, "Vaduz")
}

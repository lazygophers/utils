//go:build (lang_es || lang_all) && (country_all || country_europe || country_nl || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNetherlands.RegisterName(xlanguage.Spanish, "Países Bajos")
	dataNetherlands.RegisterOfficialName(xlanguage.Spanish, "Reino de los Países Bajos")
	dataNetherlands.RegisterCapital(xlanguage.Spanish, "Ámsterdam")
}

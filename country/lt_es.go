//go:build (lang_es || lang_all) && (country_all || country_europe || country_lt || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLithuania.RegisterName(xlanguage.Spanish, "Lituania")
	dataLithuania.RegisterOfficialName(xlanguage.Spanish, "República de Lituania")
	dataLithuania.RegisterCapital(xlanguage.Spanish, "Vilna")
}

//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLaos.RegisterName(xlanguage.Spanish, "Laos")
	dataLaos.RegisterOfficialName(xlanguage.Spanish, "República Democrática Popular Lao")
	dataLaos.RegisterCapital(xlanguage.Spanish, "Vientián")
}

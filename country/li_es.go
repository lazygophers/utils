//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLiechtenstein.RegisterName(xlanguage.Spanish, "Liechtenstein")
	dataLiechtenstein.RegisterOfficialName(xlanguage.Spanish, "Principado de Liechtenstein")
	dataLiechtenstein.RegisterCapital(xlanguage.Spanish, "Vaduz")
}

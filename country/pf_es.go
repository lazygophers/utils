//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchPolynesia.RegisterName(xlanguage.Spanish, "Polinesia Francesa")
	dataFrenchPolynesia.RegisterOfficialName(xlanguage.Spanish, "Polinesia Francesa")
	dataFrenchPolynesia.RegisterCapital(xlanguage.Spanish, "Papeete")
}

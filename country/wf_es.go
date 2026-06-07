//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataWallisAndFutuna.RegisterName(xlanguage.Spanish, "Wallis y Futuna")
	dataWallisAndFutuna.RegisterOfficialName(xlanguage.Spanish, "Territorio de las Islas Wallis y Futuna")
	dataWallisAndFutuna.RegisterCapital(xlanguage.Spanish, "Mata-Utu")
}

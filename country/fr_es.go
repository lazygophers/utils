//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrance.RegisterName(xlanguage.Spanish, "Francia")
	dataFrance.RegisterOfficialName(xlanguage.Spanish, "República Francesa")
	dataFrance.RegisterCapital(xlanguage.Spanish, "París")
}

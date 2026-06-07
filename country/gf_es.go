//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchGuiana.RegisterName(xlanguage.Spanish, "Guayana Francesa")
	dataFrenchGuiana.RegisterOfficialName(xlanguage.Spanish, "Guayana Francesa")
	dataFrenchGuiana.RegisterCapital(xlanguage.Spanish, "Cayena")
}

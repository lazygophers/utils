//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChina.RegisterName(xlanguage.French, "Chine")
	dataChina.RegisterOfficialName(xlanguage.French, "République populaire de Chine")
	dataChina.RegisterCapital(xlanguage.French, "Pékin")
}

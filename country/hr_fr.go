//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCroatia.RegisterName(xlanguage.French, "Croatie")
	dataCroatia.RegisterOfficialName(xlanguage.French, "République de Croatie")
	dataCroatia.RegisterCapital(xlanguage.French, "Zagreb")
}

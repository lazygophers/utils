//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAmericanSamoa.RegisterName(xlanguage.Spanish, "Samoa Americana")
	dataAmericanSamoa.RegisterOfficialName(xlanguage.Spanish, "Territorio de Samoa Americana")
	dataAmericanSamoa.RegisterCapital(xlanguage.Spanish, "Pago Pago")
}

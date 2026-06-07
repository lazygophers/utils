//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAmericanSamoa.RegisterName(xlanguage.French, "Samoa américaines")
	dataAmericanSamoa.RegisterOfficialName(xlanguage.French, "Samoa américaines")
	dataAmericanSamoa.RegisterCapital(xlanguage.French, "Pago Pago")
}

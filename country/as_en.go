package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAmericanSamoa.RegisterName(xlanguage.English, "American Samoa")
	dataAmericanSamoa.RegisterOfficialName(xlanguage.English, "Territory of American Samoa")
	dataAmericanSamoa.RegisterCapital(xlanguage.English, "Pago Pago")
}

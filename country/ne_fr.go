package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNiger.RegisterName(xlanguage.French, "Niger")
	dataNiger.RegisterOfficialName(xlanguage.French, "République du Niger")
	dataNiger.RegisterCapital(xlanguage.French, "Niamey")
}

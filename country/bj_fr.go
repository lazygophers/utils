package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBenin.RegisterName(xlanguage.French, "Bénin")
	dataBenin.RegisterOfficialName(xlanguage.French, "République du Bénin")
	dataBenin.RegisterCapital(xlanguage.French, "Porto-Novo")
}

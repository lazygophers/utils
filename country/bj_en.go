package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBenin.RegisterName(xlanguage.English, "Benin")
	dataBenin.RegisterOfficialName(xlanguage.English, "Republic of Benin")
	dataBenin.RegisterCapital(xlanguage.English, "Porto-Novo")
}

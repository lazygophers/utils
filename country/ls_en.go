package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLesotho.RegisterName(xlanguage.English, "Lesotho")
	dataLesotho.RegisterOfficialName(xlanguage.English, "Kingdom of Lesotho")
	dataLesotho.RegisterCapital(xlanguage.English, "Maseru")
}

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUzbekistan.RegisterName(xlanguage.Russian, "Узбекистан")
	dataUzbekistan.RegisterOfficialName(xlanguage.Russian, "Республика Узбекистан")
	dataUzbekistan.RegisterCapital(xlanguage.Russian, "Ташкент")
}

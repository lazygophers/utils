package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDominica.RegisterName(xlanguage.English, "Dominica")
	dataDominica.RegisterOfficialName(xlanguage.English, "Commonwealth of Dominica")
	dataDominica.RegisterCapital(xlanguage.English, "Roseau")
}

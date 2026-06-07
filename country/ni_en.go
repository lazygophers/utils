package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNicaragua.RegisterName(xlanguage.English, "Nicaragua")
	dataNicaragua.RegisterOfficialName(xlanguage.English, "Republic of Nicaragua")
	dataNicaragua.RegisterCapital(xlanguage.English, "Managua")
}

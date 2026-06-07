package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorway.RegisterName(xlanguage.English, "Norway")
	dataNorway.RegisterOfficialName(xlanguage.English, "Kingdom of Norway")
	dataNorway.RegisterCapital(xlanguage.English, "Oslo")
}

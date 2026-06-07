package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthAfrica.RegisterName(xlanguage.English, "South Africa")
	dataSouthAfrica.RegisterOfficialName(xlanguage.English, "Republic of South Africa")
	dataSouthAfrica.RegisterCapital(xlanguage.English, "Pretoria")
}

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSuriname.RegisterName(xlanguage.English, "Suriname")
	dataSuriname.RegisterOfficialName(xlanguage.English, "Republic of Suriname")
	dataSuriname.RegisterCapital(xlanguage.English, "Paramaribo")
}

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSeychelles.RegisterName(xlanguage.English, "Seychelles")
	dataSeychelles.RegisterOfficialName(xlanguage.English, "Republic of Seychelles")
	dataSeychelles.RegisterCapital(xlanguage.English, "Victoria")
}

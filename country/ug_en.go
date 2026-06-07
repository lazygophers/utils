package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUganda.RegisterName(xlanguage.English, "Uganda")
	dataUganda.RegisterOfficialName(xlanguage.English, "Republic of Uganda")
	dataUganda.RegisterCapital(xlanguage.English, "Kampala")
}

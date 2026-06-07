package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataOman.RegisterName(xlanguage.English, "Oman")
	dataOman.RegisterOfficialName(xlanguage.English, "Sultanate of Oman")
	dataOman.RegisterCapital(xlanguage.English, "Muscat")
}

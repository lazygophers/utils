//go:build country_all || country_asia || country_jo || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJordan.RegisterName(xlanguage.English, "Jordan")
	dataJordan.RegisterOfficialName(xlanguage.English, "Hashemite Kingdom of Jordan")
	dataJordan.RegisterCapital(xlanguage.English, "Amman")
}

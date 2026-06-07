//go:build country_all || country_asia || country_om || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataOman.RegisterName(xlanguage.English, "Oman")
	dataOman.RegisterOfficialName(xlanguage.English, "Sultanate of Oman")
	dataOman.RegisterCapital(xlanguage.English, "Muscat")
}

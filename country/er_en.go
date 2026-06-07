//go:build country_africa || country_all || country_eastern_africa || country_er

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEritrea.RegisterName(xlanguage.English, "Eritrea")
	dataEritrea.RegisterOfficialName(xlanguage.English, "State of Eritrea")
	dataEritrea.RegisterCapital(xlanguage.English, "Asmara")
}

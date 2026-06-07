//go:build country_all || country_asia || country_sa || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaudiArabia.RegisterName(xlanguage.English, "Saudi Arabia")
	dataSaudiArabia.RegisterOfficialName(xlanguage.English, "Kingdom of Saudi Arabia")
	dataSaudiArabia.RegisterCapital(xlanguage.English, "Riyadh")
}

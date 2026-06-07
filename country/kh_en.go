//go:build country_all || country_asia || country_kh || country_south_eastern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCambodia.RegisterName(xlanguage.English, "Cambodia")
	dataCambodia.RegisterOfficialName(xlanguage.English, "Kingdom of Cambodia")
	dataCambodia.RegisterCapital(xlanguage.English, "Phnom Penh")
}

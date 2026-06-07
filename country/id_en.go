//go:build country_all || country_asia || country_id || country_south_eastern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIndonesia.RegisterName(xlanguage.English, "Indonesia")
	dataIndonesia.RegisterOfficialName(xlanguage.English, "Republic of Indonesia")
	dataIndonesia.RegisterCapital(xlanguage.English, "Jakarta")
}

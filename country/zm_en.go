//go:build country_africa || country_all || country_eastern_africa || country_zm

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZambia.RegisterName(xlanguage.English, "Zambia")
	dataZambia.RegisterOfficialName(xlanguage.English, "Republic of Zambia")
	dataZambia.RegisterCapital(xlanguage.English, "Lusaka")
}

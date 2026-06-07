//go:build (lang_fr || lang_all) && (country_africa || country_all || country_eastern_africa || country_zm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZambia.RegisterName(xlanguage.French, "Zambie")
	dataZambia.RegisterOfficialName(xlanguage.French, "République de Zambie")
	dataZambia.RegisterCapital(xlanguage.French, "Lusaka")
}

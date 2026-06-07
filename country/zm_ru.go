//go:build (lang_ru || lang_all) && (country_africa || country_all || country_eastern_africa || country_zm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZambia.RegisterName(xlanguage.Russian, "Замбия")
	dataZambia.RegisterOfficialName(xlanguage.Russian, "Республика Замбия")
	dataZambia.RegisterCapital(xlanguage.Russian, "Лусака")
}

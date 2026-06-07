//go:build (lang_ru || lang_all) && (country_africa || country_all || country_eastern_africa || country_ss)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthSudan.RegisterName(xlanguage.Russian, "Южный Судан")
	dataSouthSudan.RegisterOfficialName(xlanguage.Russian, "Республика Южный Судан")
	dataSouthSudan.RegisterCapital(xlanguage.Russian, "Джуба")
}

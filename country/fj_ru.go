//go:build (lang_ru || lang_all) && (country_all || country_fj || country_melanesia || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFiji.RegisterName(xlanguage.Russian, "Фиджи")
	dataFiji.RegisterOfficialName(xlanguage.Russian, "Республика Фиджи")
	dataFiji.RegisterCapital(xlanguage.Russian, "Сува")
}

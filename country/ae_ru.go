//go:build (lang_ru || lang_all) && (country_ae || country_all || country_asia || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedArabEmirates.RegisterName(xlanguage.Russian, "Объединённые Арабские Эмираты")
	dataUnitedArabEmirates.RegisterOfficialName(xlanguage.Russian, "Объединённые Арабские Эмираты")
	dataUnitedArabEmirates.RegisterCapital(xlanguage.Russian, "Абу-Даби")
}

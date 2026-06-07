//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedArabEmirates.RegisterName(xlanguage.Russian, "Объединённые Арабские Эмираты")
	dataUnitedArabEmirates.RegisterOfficialName(xlanguage.Russian, "Объединённые Арабские Эмираты")
	dataUnitedArabEmirates.RegisterCapital(xlanguage.Russian, "Абу-Даби")
}

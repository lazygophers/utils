//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedStates.RegisterName(xlanguage.Russian, "США")
	dataUnitedStates.RegisterOfficialName(xlanguage.Russian, "Соединённые Штаты Америки")
	dataUnitedStates.RegisterCapital(xlanguage.Russian, "Вашингтон")
}

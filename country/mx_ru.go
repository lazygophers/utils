//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMexico.RegisterName(xlanguage.Russian, "Мексика")
	dataMexico.RegisterOfficialName(xlanguage.Russian, "Мексиканские Соединённые Штаты")
	dataMexico.RegisterCapital(xlanguage.Russian, "Мехико")
}

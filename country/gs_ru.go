//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthGeorgiaAndSouthSandwich.RegisterName(xlanguage.Russian, "Южная Георгия и Южные Сандвичевы Острова")
	dataSouthGeorgiaAndSouthSandwich.RegisterOfficialName(xlanguage.Russian, "Южная Георгия и Южные Сандвичевы Острова")
	dataSouthGeorgiaAndSouthSandwich.RegisterCapital(xlanguage.Russian, "Кинг-Эдуард-Пойнт")
}

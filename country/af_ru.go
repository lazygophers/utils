//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAfghanistan.RegisterName(xlanguage.Russian, "Афганистан")
	dataAfghanistan.RegisterOfficialName(xlanguage.Russian, "Исламский Эмират Афганистан")
	dataAfghanistan.RegisterCapital(xlanguage.Russian, "Кабул")
}

//go:build (lang_ru || lang_all) && (country_af || country_all || country_asia || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAfghanistan.RegisterName(xlanguage.Russian, "Афганистан")
	dataAfghanistan.RegisterOfficialName(xlanguage.Russian, "Исламский Эмират Афганистан")
	dataAfghanistan.RegisterCapital(xlanguage.Russian, "Кабул")
}

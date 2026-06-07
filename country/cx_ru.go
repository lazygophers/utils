//go:build (lang_ru || lang_all) && (country_all || country_australia_and_new_zealand || country_cx || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChristmasIsland.RegisterName(xlanguage.Russian, "Остров Рождества")
	dataChristmasIsland.RegisterOfficialName(xlanguage.Russian, "Территория Остров Рождества")
	dataChristmasIsland.RegisterCapital(xlanguage.Russian, "Флайинг-Фиш-Ков")
}

//go:build (lang_ru || lang_all) && (country_africa || country_all || country_eh || country_northern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataWesternSahara.RegisterName(xlanguage.Russian, "Западная Сахара")
	dataWesternSahara.RegisterOfficialName(xlanguage.Russian, "Сахарская Арабская Демократическая Республика")
	dataWesternSahara.RegisterCapital(xlanguage.Russian, "Эль-Аюн")
}

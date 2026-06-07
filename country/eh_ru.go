//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataWesternSahara.RegisterName(xlanguage.Russian, "Западная Сахара")
	dataWesternSahara.RegisterOfficialName(xlanguage.Russian, "Сахарская Арабская Демократическая Республика")
	dataWesternSahara.RegisterCapital(xlanguage.Russian, "Эль-Аюн")
}

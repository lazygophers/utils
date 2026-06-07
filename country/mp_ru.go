//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthernMarianaIslands.RegisterName(xlanguage.Russian, "Северные Марианские Острова")
	dataNorthernMarianaIslands.RegisterOfficialName(xlanguage.Russian, "Содружество Северных Марианских Островов")
	dataNorthernMarianaIslands.RegisterCapital(xlanguage.Russian, "Сайпан")
}

//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFalklandIslands.RegisterName(xlanguage.Russian, "Фолклендские острова")
	dataFalklandIslands.RegisterOfficialName(xlanguage.Russian, "Фолклендские острова")
	dataFalklandIslands.RegisterCapital(xlanguage.Russian, "Стэнли")
}

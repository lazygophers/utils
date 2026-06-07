//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSwitzerland.RegisterName(xlanguage.Russian, "Швейцария")
	dataSwitzerland.RegisterOfficialName(xlanguage.Russian, "Швейцарская Конфедерация")
	dataSwitzerland.RegisterCapital(xlanguage.Russian, "Берн")
}

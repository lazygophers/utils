//go:build (lang_ru || lang_all) && (country_all || country_ch || country_europe || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSwitzerland.RegisterName(xlanguage.Russian, "Швейцария")
	dataSwitzerland.RegisterOfficialName(xlanguage.Russian, "Швейцарская Конфедерация")
	dataSwitzerland.RegisterCapital(xlanguage.Russian, "Берн")
}

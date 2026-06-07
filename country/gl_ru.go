//go:build (lang_ru || lang_all) && (country_all || country_americas || country_gl || country_northern_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGreenland.RegisterName(xlanguage.Russian, "Гренландия")
	dataGreenland.RegisterOfficialName(xlanguage.Russian, "Гренландия")
	dataGreenland.RegisterCapital(xlanguage.Russian, "Нуук")
}

//go:build (lang_ru || lang_all) && (country_all || country_es || country_europe || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSpain.RegisterName(xlanguage.Russian, "Испания")
	dataSpain.RegisterOfficialName(xlanguage.Russian, "Королевство Испания")
	dataSpain.RegisterCapital(xlanguage.Russian, "Мадрид")
}

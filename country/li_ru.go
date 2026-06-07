//go:build (lang_ru || lang_all) && (country_all || country_europe || country_li || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLiechtenstein.RegisterName(xlanguage.Russian, "Лихтенштейн")
	dataLiechtenstein.RegisterOfficialName(xlanguage.Russian, "Княжество Лихтенштейн")
	dataLiechtenstein.RegisterCapital(xlanguage.Russian, "Вадуц")
}

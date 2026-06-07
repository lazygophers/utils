//go:build (lang_ru || lang_all) && (country_all || country_ee || country_europe || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEstonia.RegisterName(xlanguage.Russian, "Эстония")
	dataEstonia.RegisterOfficialName(xlanguage.Russian, "Эстонская Республика")
	dataEstonia.RegisterCapital(xlanguage.Russian, "Таллин")
}

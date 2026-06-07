//go:build (lang_ru || lang_all) && (country_all || country_europe || country_si || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSlovenia.RegisterName(xlanguage.Russian, "Словения")
	dataSlovenia.RegisterOfficialName(xlanguage.Russian, "Республика Словения")
	dataSlovenia.RegisterCapital(xlanguage.Russian, "Любляна")
}

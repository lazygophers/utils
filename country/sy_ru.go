//go:build (lang_ru || lang_all) && (country_all || country_asia || country_sy || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSyria.RegisterName(xlanguage.Russian, "Сирия")
	dataSyria.RegisterOfficialName(xlanguage.Russian, "Сирийская Арабская Республика")
	dataSyria.RegisterCapital(xlanguage.Russian, "Дамаск")
}

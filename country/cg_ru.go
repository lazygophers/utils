//go:build (lang_ru || lang_all) && (country_africa || country_all || country_cg || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCongo.RegisterName(xlanguage.Russian, "Республика Конго")
	dataCongo.RegisterOfficialName(xlanguage.Russian, "Республика Конго")
	dataCongo.RegisterCapital(xlanguage.Russian, "Браззавиль")
}

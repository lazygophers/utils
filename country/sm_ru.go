//go:build (lang_ru || lang_all) && (country_all || country_europe || country_sm || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSanMarino.RegisterName(xlanguage.Russian, "Сан-Марино")
	dataSanMarino.RegisterOfficialName(xlanguage.Russian, "Республика Сан-Марино")
	dataSanMarino.RegisterCapital(xlanguage.Russian, "Сан-Марино")
}

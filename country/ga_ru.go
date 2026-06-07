//go:build (lang_ru || lang_all) && (country_africa || country_all || country_ga || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGabon.RegisterName(xlanguage.Russian, "Габон")
	dataGabon.RegisterOfficialName(xlanguage.Russian, "Габонская Республика")
	dataGabon.RegisterCapital(xlanguage.Russian, "Либревиль")
}

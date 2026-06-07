//go:build (lang_ru || lang_all) && (country_africa || country_all || country_eastern_africa || country_mw)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalawi.RegisterName(xlanguage.Russian, "Малави")
	dataMalawi.RegisterOfficialName(xlanguage.Russian, "Республика Малави")
	dataMalawi.RegisterCapital(xlanguage.Russian, "Лилонгве")
}

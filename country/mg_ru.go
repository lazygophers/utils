//go:build (lang_ru || lang_all) && (country_africa || country_all || country_eastern_africa || country_mg)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMadagascar.RegisterName(xlanguage.Russian, "Мадагаскар")
	dataMadagascar.RegisterOfficialName(xlanguage.Russian, "Республика Мадагаскар")
	dataMadagascar.RegisterCapital(xlanguage.Russian, "Антананариву")
}

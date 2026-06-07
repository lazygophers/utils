//go:build (lang_ru || lang_all) && (country_all || country_oceania || country_pf || country_polynesia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchPolynesia.RegisterName(xlanguage.Russian, "Французская Полинезия")
	dataFrenchPolynesia.RegisterOfficialName(xlanguage.Russian, "Французская Полинезия")
	dataFrenchPolynesia.RegisterCapital(xlanguage.Russian, "Папеэте")
}

//go:build (lang_ru || lang_all) && (country_africa || country_all || country_gq || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEquatorialGuinea.RegisterName(xlanguage.Russian, "Экваториальная Гвинея")
	dataEquatorialGuinea.RegisterOfficialName(xlanguage.Russian, "Республика Экваториальная Гвинея")
	dataEquatorialGuinea.RegisterCapital(xlanguage.Russian, "Малабо")
}

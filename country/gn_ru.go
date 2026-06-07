//go:build (lang_ru || lang_all) && (country_africa || country_all || country_gn || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuinea.RegisterName(xlanguage.Russian, "Гвинея")
	dataGuinea.RegisterOfficialName(xlanguage.Russian, "Гвинейская Республика")
	dataGuinea.RegisterCapital(xlanguage.Russian, "Конакри")
}

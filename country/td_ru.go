//go:build (lang_ru || lang_all) && (country_africa || country_all || country_middle_africa || country_td)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChad.RegisterName(xlanguage.Russian, "Чад")
	dataChad.RegisterOfficialName(xlanguage.Russian, "Республика Чад")
	dataChad.RegisterCapital(xlanguage.Russian, "Нджамена")
}

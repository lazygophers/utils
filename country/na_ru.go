//go:build (lang_ru || lang_all) && (country_africa || country_all || country_na || country_southern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNamibia.RegisterName(xlanguage.Russian, "Намибия")
	dataNamibia.RegisterOfficialName(xlanguage.Russian, "Республика Намибия")
	dataNamibia.RegisterCapital(xlanguage.Russian, "Виндхук")
}

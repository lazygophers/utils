//go:build (lang_ru || lang_all) && (country_africa || country_all || country_northern_africa || country_tn)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTunisia.RegisterName(xlanguage.Russian, "Тунис")
	dataTunisia.RegisterOfficialName(xlanguage.Russian, "Тунисская Республика")
	dataTunisia.RegisterCapital(xlanguage.Russian, "Тунис")
}

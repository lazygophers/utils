//go:build (lang_ru || lang_all) && (country_africa || country_all || country_eastern_africa || country_ke)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKenya.RegisterName(xlanguage.Russian, "Кения")
	dataKenya.RegisterOfficialName(xlanguage.Russian, "Республика Кения")
	dataKenya.RegisterCapital(xlanguage.Russian, "Найроби")
}

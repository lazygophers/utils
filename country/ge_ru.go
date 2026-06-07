//go:build (lang_ru || lang_all) && (country_all || country_asia || country_ge || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGeorgia.RegisterName(xlanguage.Russian, "Грузия")
	dataGeorgia.RegisterOfficialName(xlanguage.Russian, "Грузия")
	dataGeorgia.RegisterCapital(xlanguage.Russian, "Тбилиси")
}

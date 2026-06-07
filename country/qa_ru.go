//go:build (lang_ru || lang_all) && (country_all || country_asia || country_qa || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataQatar.RegisterName(xlanguage.Russian, "Катар")
	dataQatar.RegisterOfficialName(xlanguage.Russian, "Государство Катар")
	dataQatar.RegisterCapital(xlanguage.Russian, "Доха")
}

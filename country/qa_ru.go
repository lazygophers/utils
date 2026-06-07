//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataQatar.RegisterName(xlanguage.Russian, "Катар")
	dataQatar.RegisterOfficialName(xlanguage.Russian, "Государство Катар")
	dataQatar.RegisterCapital(xlanguage.Russian, "Доха")
}

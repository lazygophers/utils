//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLibya.RegisterName(xlanguage.Russian, "Ливия")
	dataLibya.RegisterOfficialName(xlanguage.Russian, "Государство Ливия")
	dataLibya.RegisterCapital(xlanguage.Russian, "Триполи")
}

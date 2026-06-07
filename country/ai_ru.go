//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAnguilla.RegisterName(xlanguage.Russian, "Ангилья")
	dataAnguilla.RegisterOfficialName(xlanguage.Russian, "Ангилья")
	dataAnguilla.RegisterCapital(xlanguage.Russian, "Валли")
}

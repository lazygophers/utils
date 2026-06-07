//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEswatini.RegisterName(xlanguage.Russian, "Эсватини")
	dataEswatini.RegisterOfficialName(xlanguage.Russian, "Королевство Эсватини")
	dataEswatini.RegisterCapital(xlanguage.Russian, "Мбабане")
}

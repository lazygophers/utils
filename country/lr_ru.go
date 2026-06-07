//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLiberia.RegisterName(xlanguage.Russian, "Либерия")
	dataLiberia.RegisterOfficialName(xlanguage.Russian, "Республика Либерия")
	dataLiberia.RegisterCapital(xlanguage.Russian, "Монровия")
}

//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrazil.RegisterName(xlanguage.French, "Brésil")
	dataBrazil.RegisterOfficialName(xlanguage.French, "République fédérative du Brésil")
	dataBrazil.RegisterCapital(xlanguage.French, "Brasilia")
}

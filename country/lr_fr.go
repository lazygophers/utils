//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLiberia.RegisterName(xlanguage.French, "Liberia")
	dataLiberia.RegisterOfficialName(xlanguage.French, "République du Liberia")
	dataLiberia.RegisterCapital(xlanguage.French, "Monrovia")
}

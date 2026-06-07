//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuineaBissau.RegisterName(xlanguage.French, "Guinée-Bissau")
	dataGuineaBissau.RegisterOfficialName(xlanguage.French, "République de Guinée-Bissau")
	dataGuineaBissau.RegisterCapital(xlanguage.French, "Bissau")
}

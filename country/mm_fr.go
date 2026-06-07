//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMyanmar.RegisterName(xlanguage.French, "Birmanie")
	dataMyanmar.RegisterOfficialName(xlanguage.French, "République de l'Union du Myanmar")
	dataMyanmar.RegisterCapital(xlanguage.French, "Naypyidaw")
}

//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAnguilla.RegisterName(xlanguage.French, "Anguilla")
	dataAnguilla.RegisterOfficialName(xlanguage.French, "Anguilla")
	dataAnguilla.RegisterCapital(xlanguage.French, "The Valley")
}

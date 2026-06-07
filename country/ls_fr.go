//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLesotho.RegisterName(xlanguage.French, "Lesotho")
	dataLesotho.RegisterOfficialName(xlanguage.French, "Royaume du Lesotho")
	dataLesotho.RegisterCapital(xlanguage.French, "Maseru")
}

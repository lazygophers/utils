//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSriLanka.RegisterName(xlanguage.French, "Sri Lanka")
	dataSriLanka.RegisterOfficialName(xlanguage.French, "République socialiste démocratique du Sri Lanka")
	dataSriLanka.RegisterCapital(xlanguage.French, "Sri Jayawardenapura Kotte")
}

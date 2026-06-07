//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCzechia.RegisterName(xlanguage.French, "Tchéquie")
	dataCzechia.RegisterOfficialName(xlanguage.French, "République tchèque")
	dataCzechia.RegisterCapital(xlanguage.French, "Prague")
}

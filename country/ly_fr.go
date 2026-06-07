//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLibya.RegisterName(xlanguage.French, "Libye")
	dataLibya.RegisterOfficialName(xlanguage.French, "État de Libye")
	dataLibya.RegisterCapital(xlanguage.French, "Tripoli")
}

//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBhutan.RegisterName(xlanguage.French, "Bhoutan")
	dataBhutan.RegisterOfficialName(xlanguage.French, "Royaume du Bhoutan")
	dataBhutan.RegisterCapital(xlanguage.French, "Thimphou")
}

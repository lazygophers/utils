//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPanama.RegisterName(xlanguage.French, "Panama")
	dataPanama.RegisterOfficialName(xlanguage.French, "République du Panama")
	dataPanama.RegisterCapital(xlanguage.French, "Panama")
}

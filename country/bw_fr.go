//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBotswana.RegisterName(xlanguage.French, "Botswana")
	dataBotswana.RegisterOfficialName(xlanguage.French, "République du Botswana")
	dataBotswana.RegisterCapital(xlanguage.French, "Gaborone")
}

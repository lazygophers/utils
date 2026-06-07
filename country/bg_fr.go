//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBulgaria.RegisterName(xlanguage.French, "Bulgarie")
	dataBulgaria.RegisterOfficialName(xlanguage.French, "République de Bulgarie")
	dataBulgaria.RegisterCapital(xlanguage.French, "Sofia")
}

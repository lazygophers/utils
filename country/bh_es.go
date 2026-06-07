//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBahrain.RegisterName(xlanguage.Spanish, "Baréin")
	dataBahrain.RegisterOfficialName(xlanguage.Spanish, "Reino de Baréin")
	dataBahrain.RegisterCapital(xlanguage.Spanish, "Manama")
}

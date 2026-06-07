//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBahrain.RegisterName(xlanguage.French, "Bahreïn")
	dataBahrain.RegisterOfficialName(xlanguage.French, "Royaume de Bahreïn")
	dataBahrain.RegisterCapital(xlanguage.French, "Manama")
}

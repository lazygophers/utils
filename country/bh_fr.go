//go:build (lang_fr || lang_all) && (country_all || country_asia || country_bh || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBahrain.RegisterName(xlanguage.French, "Bahreïn")
	dataBahrain.RegisterOfficialName(xlanguage.French, "Royaume de Bahreïn")
	dataBahrain.RegisterCapital(xlanguage.French, "Manama")
}

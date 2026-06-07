//go:build country_all || country_asia || country_bh || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBahrain.RegisterName(xlanguage.English, "Bahrain")
	dataBahrain.RegisterOfficialName(xlanguage.English, "Kingdom of Bahrain")
	dataBahrain.RegisterCapital(xlanguage.English, "Manama")
}

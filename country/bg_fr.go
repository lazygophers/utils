//go:build (lang_fr || lang_all) && (country_all || country_bg || country_eastern_europe || country_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBulgaria.RegisterName(xlanguage.French, "Bulgarie")
	dataBulgaria.RegisterOfficialName(xlanguage.French, "République de Bulgarie")
	dataBulgaria.RegisterCapital(xlanguage.French, "Sofia")
}

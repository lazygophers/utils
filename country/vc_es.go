//go:build (lang_es || lang_all) && (country_all || country_americas || country_caribbean || country_vc)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintVincentAndGrenadines.RegisterName(xlanguage.Spanish, "San Vicente y las Granadinas")
	dataSaintVincentAndGrenadines.RegisterOfficialName(xlanguage.Spanish, "San Vicente y las Granadinas")
	dataSaintVincentAndGrenadines.RegisterCapital(xlanguage.Spanish, "Kingstown")
}

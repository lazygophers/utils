//go:build (lang_es || lang_all) && (country_all || country_americas || country_bs || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBahamas.RegisterName(xlanguage.Spanish, "Bahamas")
	dataBahamas.RegisterOfficialName(xlanguage.Spanish, "Mancomunidad de las Bahamas")
	dataBahamas.RegisterCapital(xlanguage.Spanish, "Nasáu")
}

//go:build (lang_fr || lang_all) && (country_all || country_americas || country_bs || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBahamas.RegisterName(xlanguage.French, "Bahamas")
	dataBahamas.RegisterOfficialName(xlanguage.French, "Commonwealth des Bahamas")
	dataBahamas.RegisterCapital(xlanguage.French, "Nassau")
}

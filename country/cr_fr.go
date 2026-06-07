//go:build (lang_fr || lang_all) && (country_all || country_americas || country_central_america || country_cr)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCostaRica.RegisterName(xlanguage.French, "Costa Rica")
	dataCostaRica.RegisterOfficialName(xlanguage.French, "République du Costa Rica")
	dataCostaRica.RegisterCapital(xlanguage.French, "San José")
}

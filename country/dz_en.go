//go:build country_africa || country_all || country_dz || country_northern_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlgeria.RegisterName(xlanguage.English, "Algeria")
	dataAlgeria.RegisterOfficialName(xlanguage.English, "People's Democratic Republic of Algeria")
	dataAlgeria.RegisterCapital(xlanguage.English, "Algiers")
}

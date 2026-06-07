//go:build country_all || country_asia || country_central_asia || country_kz

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKazakhstan.RegisterName(xlanguage.English, "Kazakhstan")
	dataKazakhstan.RegisterOfficialName(xlanguage.English, "Republic of Kazakhstan")
	dataKazakhstan.RegisterCapital(xlanguage.English, "Astana")
}

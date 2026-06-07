//go:build country_all || country_americas || country_central_america || country_cr

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCostaRica.RegisterName(xlanguage.English, "Costa Rica")
	dataCostaRica.RegisterOfficialName(xlanguage.English, "Republic of Costa Rica")
	dataCostaRica.RegisterCapital(xlanguage.English, "San Jose")
}

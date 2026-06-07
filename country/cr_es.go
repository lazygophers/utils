//go:build country_all || country_americas || country_central_america || country_cr

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCostaRica.RegisterName(xlanguage.Spanish, "Costa Rica")
	dataCostaRica.RegisterOfficialName(xlanguage.Spanish, "República de Costa Rica")
	dataCostaRica.RegisterCapital(xlanguage.Spanish, "San José")
}

//go:build country_all || country_asia || country_central_asia || country_tj

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTajikistan.RegisterName(xlanguage.Russian, "Таджикистан")
	dataTajikistan.RegisterOfficialName(xlanguage.Russian, "Республика Таджикистан")
	dataTajikistan.RegisterCapital(xlanguage.Russian, "Душанбе")
}

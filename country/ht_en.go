//go:build country_all || country_americas || country_caribbean || country_ht

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHaiti.RegisterName(xlanguage.English, "Haiti")
	dataHaiti.RegisterOfficialName(xlanguage.English, "Republic of Haiti")
	dataHaiti.RegisterCapital(xlanguage.English, "Port-au-Prince")
}

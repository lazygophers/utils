//go:build country_africa || country_all || country_southern_africa || country_za

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthAfrica.RegisterName(xlanguage.English, "South Africa")
	dataSouthAfrica.RegisterOfficialName(xlanguage.English, "Republic of South Africa")
	dataSouthAfrica.RegisterCapital(xlanguage.English, "Pretoria")
}

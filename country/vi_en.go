//go:build country_all || country_americas || country_caribbean || country_vi

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUsVirginIslands.RegisterName(xlanguage.English, "United States Virgin Islands")
	dataUsVirginIslands.RegisterOfficialName(xlanguage.English, "Virgin Islands of the United States")
	dataUsVirginIslands.RegisterCapital(xlanguage.English, "Charlotte Amalie")
}

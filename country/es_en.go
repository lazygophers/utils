//go:build country_all || country_es || country_europe || country_southern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSpain.RegisterName(xlanguage.English, "Spain")
	dataSpain.RegisterOfficialName(xlanguage.English, "Kingdom of Spain")
	dataSpain.RegisterCapital(xlanguage.English, "Madrid")
}

//go:build country_all || country_europe || country_lt || country_northern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLithuania.RegisterName(xlanguage.English, "Lithuania")
	dataLithuania.RegisterOfficialName(xlanguage.English, "Republic of Lithuania")
	dataLithuania.RegisterCapital(xlanguage.English, "Vilnius")
}

//go:build country_all || country_europe || country_ie || country_northern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIreland.RegisterName(xlanguage.English, "Ireland")
	dataIreland.RegisterOfficialName(xlanguage.English, "Republic of Ireland")
	dataIreland.RegisterCapital(xlanguage.English, "Dublin")
}

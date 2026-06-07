//go:build country_all || country_europe || country_it || country_southern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataItaly.RegisterName(xlanguage.English, "Italy")
	dataItaly.RegisterOfficialName(xlanguage.English, "Italian Republic")
	dataItaly.RegisterCapital(xlanguage.English, "Rome")
}

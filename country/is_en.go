//go:build country_all || country_europe || country_is || country_northern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIceland.RegisterName(xlanguage.English, "Iceland")
	dataIceland.RegisterOfficialName(xlanguage.English, "Iceland")
	dataIceland.RegisterCapital(xlanguage.English, "Reykjavik")
}

//go:build country_all || country_ax || country_europe || country_northern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlandIslands.RegisterName(xlanguage.English, "Aland Islands")
	dataAlandIslands.RegisterOfficialName(xlanguage.English, "Aland Islands")
	dataAlandIslands.RegisterCapital(xlanguage.English, "Mariehamn")
}

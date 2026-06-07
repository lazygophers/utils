//go:build country_all || country_ch || country_europe || country_western_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSwitzerland.RegisterName(xlanguage.English, "Switzerland")
	dataSwitzerland.RegisterOfficialName(xlanguage.English, "Swiss Confederation")
	dataSwitzerland.RegisterCapital(xlanguage.English, "Bern")
}

//go:build country_all || country_americas || country_caribbean || country_vg

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishVirginIslands.RegisterName(xlanguage.English, "British Virgin Islands")
	dataBritishVirginIslands.RegisterOfficialName(xlanguage.English, "Virgin Islands")
	dataBritishVirginIslands.RegisterCapital(xlanguage.English, "Road Town")
}

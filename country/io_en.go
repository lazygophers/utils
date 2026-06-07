//go:build country_all || country_asia || country_eastern_africa || country_io

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishIndianOceanTerritory.RegisterName(xlanguage.English, "British Indian Ocean Territory")
	dataBritishIndianOceanTerritory.RegisterOfficialName(xlanguage.English, "British Indian Ocean Territory")
	dataBritishIndianOceanTerritory.RegisterCapital(xlanguage.English, "Diego Garcia")
}

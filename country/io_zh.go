//go:build country_all || country_asia || country_eastern_africa || country_io

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishIndianOceanTerritory.RegisterName(xlanguage.Chinese, "英属印度洋领地")
	dataBritishIndianOceanTerritory.RegisterOfficialName(xlanguage.Chinese, "英属印度洋领地")
	dataBritishIndianOceanTerritory.RegisterCapital(xlanguage.Chinese, "迪戈加西亚岛")
}

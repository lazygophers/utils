//go:build country_all || country_asia || country_bn || country_south_eastern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrunei.RegisterName(xlanguage.English, "Brunei")
	dataBrunei.RegisterOfficialName(xlanguage.English, "Nation of Brunei, the Abode of Peace")
	dataBrunei.RegisterCapital(xlanguage.English, "Bandar Seri Begawan")
}

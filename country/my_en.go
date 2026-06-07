//go:build country_all || country_asia || country_my || country_south_eastern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalaysia.RegisterName(xlanguage.English, "Malaysia")
	dataMalaysia.RegisterOfficialName(xlanguage.English, "Malaysia")
	dataMalaysia.RegisterCapital(xlanguage.English, "Kuala Lumpur")
}

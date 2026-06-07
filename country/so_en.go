//go:build country_africa || country_all || country_eastern_africa || country_so

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSomalia.RegisterName(xlanguage.English, "Somalia")
	dataSomalia.RegisterOfficialName(xlanguage.English, "Federal Republic of Somalia")
	dataSomalia.RegisterCapital(xlanguage.English, "Mogadishu")
}

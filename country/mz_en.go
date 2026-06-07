//go:build country_africa || country_all || country_eastern_africa || country_mz

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMozambique.RegisterName(xlanguage.English, "Mozambique")
	dataMozambique.RegisterOfficialName(xlanguage.English, "Republic of Mozambique")
	dataMozambique.RegisterCapital(xlanguage.English, "Maputo")
}

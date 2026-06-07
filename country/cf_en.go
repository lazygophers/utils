//go:build country_africa || country_all || country_cf || country_middle_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCentralAfricanRepublic.RegisterName(xlanguage.English, "Central African Republic")
	dataCentralAfricanRepublic.RegisterOfficialName(xlanguage.English, "Central African Republic")
	dataCentralAfricanRepublic.RegisterCapital(xlanguage.English, "Bangui")
}

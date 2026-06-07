//go:build country_africa || country_all || country_eastern_africa || country_mu

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMauritius.RegisterName(xlanguage.English, "Mauritius")
	dataMauritius.RegisterOfficialName(xlanguage.English, "Republic of Mauritius")
	dataMauritius.RegisterCapital(xlanguage.English, "Port Louis")
}

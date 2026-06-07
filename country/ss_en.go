//go:build country_africa || country_all || country_eastern_africa || country_ss

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthSudan.RegisterName(xlanguage.English, "South Sudan")
	dataSouthSudan.RegisterOfficialName(xlanguage.English, "Republic of South Sudan")
	dataSouthSudan.RegisterCapital(xlanguage.English, "Juba")
}

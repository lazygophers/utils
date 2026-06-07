//go:build country_africa || country_all || country_eastern_africa || country_sc

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSeychelles.RegisterName(xlanguage.English, "Seychelles")
	dataSeychelles.RegisterOfficialName(xlanguage.English, "Republic of Seychelles")
	dataSeychelles.RegisterCapital(xlanguage.English, "Victoria")
}

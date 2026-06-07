//go:build country_africa || country_all || country_middle_africa || country_td

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChad.RegisterName(xlanguage.English, "Chad")
	dataChad.RegisterOfficialName(xlanguage.English, "Republic of Chad")
	dataChad.RegisterCapital(xlanguage.English, "N'Djamena")
}

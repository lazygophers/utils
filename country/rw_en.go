//go:build country_africa || country_all || country_eastern_africa || country_rw

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRwanda.RegisterName(xlanguage.English, "Rwanda")
	dataRwanda.RegisterOfficialName(xlanguage.English, "Republic of Rwanda")
	dataRwanda.RegisterCapital(xlanguage.English, "Kigali")
}

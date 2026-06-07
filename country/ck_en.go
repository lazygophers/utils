//go:build country_all || country_ck || country_oceania || country_polynesia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCookIslands.RegisterName(xlanguage.English, "Cook Islands")
	dataCookIslands.RegisterOfficialName(xlanguage.English, "Cook Islands")
	dataCookIslands.RegisterCapital(xlanguage.English, "Avarua")
}

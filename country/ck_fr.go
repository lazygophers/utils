//go:build (lang_fr || lang_all) && (country_all || country_ck || country_oceania || country_polynesia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCookIslands.RegisterName(xlanguage.French, "Îles Cook")
	dataCookIslands.RegisterOfficialName(xlanguage.French, "Îles Cook")
	dataCookIslands.RegisterCapital(xlanguage.French, "Avarua")
}

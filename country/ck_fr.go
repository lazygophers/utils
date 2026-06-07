//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCookIslands.RegisterName(xlanguage.French, "Îles Cook")
	dataCookIslands.RegisterOfficialName(xlanguage.French, "Îles Cook")
	dataCookIslands.RegisterCapital(xlanguage.French, "Avarua")
}

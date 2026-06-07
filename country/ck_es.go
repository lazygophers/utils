//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCookIslands.RegisterName(xlanguage.Spanish, "Islas Cook")
	dataCookIslands.RegisterOfficialName(xlanguage.Spanish, "Islas Cook")
	dataCookIslands.RegisterCapital(xlanguage.Spanish, "Avarua")
}

//go:build country_all || country_asia || country_ge || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGeorgia.RegisterName(xlanguage.English, "Georgia")
	dataGeorgia.RegisterOfficialName(xlanguage.English, "Georgia")
	dataGeorgia.RegisterCapital(xlanguage.English, "Tbilisi")
}

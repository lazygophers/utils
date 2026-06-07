//go:build (lang_es || lang_all) && (country_all || country_asia || country_ge || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGeorgia.RegisterName(xlanguage.Spanish, "Georgia")
	dataGeorgia.RegisterOfficialName(xlanguage.Spanish, "Georgia")
	dataGeorgia.RegisterCapital(xlanguage.Spanish, "Tiflis")
}

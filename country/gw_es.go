//go:build (lang_es || lang_all) && (country_africa || country_all || country_gw || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuineaBissau.RegisterName(xlanguage.Spanish, "Guinea-Bisáu")
	dataGuineaBissau.RegisterOfficialName(xlanguage.Spanish, "República de Guinea-Bisáu")
	dataGuineaBissau.RegisterCapital(xlanguage.Spanish, "Bisáu")
}

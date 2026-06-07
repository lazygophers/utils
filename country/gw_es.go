//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuineaBissau.RegisterName(xlanguage.Spanish, "Guinea-Bisáu")
	dataGuineaBissau.RegisterOfficialName(xlanguage.Spanish, "República de Guinea-Bisáu")
	dataGuineaBissau.RegisterCapital(xlanguage.Spanish, "Bisáu")
}

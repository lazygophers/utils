//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArgentina.RegisterName(xlanguage.French, "Argentine")
	dataArgentina.RegisterOfficialName(xlanguage.French, "République argentine")
	dataArgentina.RegisterCapital(xlanguage.French, "Buenos Aires")
}

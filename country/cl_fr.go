//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChile.RegisterName(xlanguage.French, "Chili")
	dataChile.RegisterOfficialName(xlanguage.French, "République du Chili")
	dataChile.RegisterCapital(xlanguage.French, "Santiago")
}

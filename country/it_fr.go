//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataItaly.RegisterName(xlanguage.French, "Italie")
	dataItaly.RegisterOfficialName(xlanguage.French, "République italienne")
	dataItaly.RegisterCapital(xlanguage.French, "Rome")
}

//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSyria.RegisterName(xlanguage.French, "Syrie")
	dataSyria.RegisterOfficialName(xlanguage.French, "République arabe syrienne")
	dataSyria.RegisterCapital(xlanguage.French, "Damas")
}

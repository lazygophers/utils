//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTajikistan.RegisterName(xlanguage.French, "Tadjikistan")
	dataTajikistan.RegisterOfficialName(xlanguage.French, "République du Tadjikistan")
	dataTajikistan.RegisterCapital(xlanguage.French, "Douchanbé")
}

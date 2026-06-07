//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataYemen.RegisterName(xlanguage.French, "Yémen")
	dataYemen.RegisterOfficialName(xlanguage.French, "République du Yémen")
	dataYemen.RegisterCapital(xlanguage.French, "Sanaa")
}

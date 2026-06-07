//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKenya.RegisterName(xlanguage.French, "Kenya")
	dataKenya.RegisterOfficialName(xlanguage.French, "République du Kenya")
	dataKenya.RegisterCapital(xlanguage.French, "Nairobi")
}

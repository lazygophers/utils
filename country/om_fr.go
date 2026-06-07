//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataOman.RegisterName(xlanguage.French, "Oman")
	dataOman.RegisterOfficialName(xlanguage.French, "Sultanat d'Oman")
	dataOman.RegisterCapital(xlanguage.French, "Mascate")
}

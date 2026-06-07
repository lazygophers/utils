//go:build (lang_fr || lang_all) && (country_all || country_asia || country_om || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataOman.RegisterName(xlanguage.French, "Oman")
	dataOman.RegisterOfficialName(xlanguage.French, "Sultanat d'Oman")
	dataOman.RegisterCapital(xlanguage.French, "Mascate")
}

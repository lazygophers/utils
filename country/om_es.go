//go:build (lang_es || lang_all) && (country_all || country_asia || country_om || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataOman.RegisterName(xlanguage.Spanish, "Omán")
	dataOman.RegisterOfficialName(xlanguage.Spanish, "Sultanato de Omán")
	dataOman.RegisterCapital(xlanguage.Spanish, "Mascate")
}

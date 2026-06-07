//go:build country_all || country_asia || country_qa || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataQatar.RegisterName(xlanguage.Arabic, "قطر")
	dataQatar.RegisterOfficialName(xlanguage.Arabic, "دولة قطر")
	dataQatar.RegisterCapital(xlanguage.Arabic, "الدوحة")
}

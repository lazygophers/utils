//go:build country_all || country_asia || country_om || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataOman.RegisterName(xlanguage.Arabic, "عُمان")
	dataOman.RegisterOfficialName(xlanguage.Arabic, "سلطنة عُمان")
	dataOman.RegisterCapital(xlanguage.Arabic, "مسقط")
}

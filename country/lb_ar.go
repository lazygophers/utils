//go:build country_all || country_asia || country_lb || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLebanon.RegisterName(xlanguage.Arabic, "لبنان")
	dataLebanon.RegisterOfficialName(xlanguage.Arabic, "الجمهورية اللبنانية")
	dataLebanon.RegisterCapital(xlanguage.Arabic, "بيروت")
}

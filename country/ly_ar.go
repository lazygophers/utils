//go:build country_africa || country_all || country_ly || country_northern_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLibya.RegisterName(xlanguage.Arabic, "ليبيا")
	dataLibya.RegisterOfficialName(xlanguage.Arabic, "دولة ليبيا")
	dataLibya.RegisterCapital(xlanguage.Arabic, "طرابلس")
}

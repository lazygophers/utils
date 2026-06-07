//go:build (lang_ko || lang_all) && (country_africa || country_all || country_dj || country_eastern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDjibouti.RegisterName(xlanguage.Korean, "지부티")
	dataDjibouti.RegisterOfficialName(xlanguage.Korean, "지부티 공화국")
	dataDjibouti.RegisterCapital(xlanguage.Korean, "지부티")
}

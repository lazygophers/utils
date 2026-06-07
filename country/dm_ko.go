//go:build (lang_ko || lang_all) && (country_all || country_americas || country_caribbean || country_dm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDominica.RegisterName(xlanguage.Korean, "도미니카 연방")
	dataDominica.RegisterOfficialName(xlanguage.Korean, "도미니카 연방")
	dataDominica.RegisterCapital(xlanguage.Korean, "로조")
}

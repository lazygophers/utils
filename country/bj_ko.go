//go:build (lang_ko || lang_all) && (country_africa || country_all || country_bj || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBenin.RegisterName(xlanguage.Korean, "베냉")
	dataBenin.RegisterOfficialName(xlanguage.Korean, "베냉 공화국")
	dataBenin.RegisterCapital(xlanguage.Korean, "포르토노보")
}

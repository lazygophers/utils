//go:build (lang_ko || lang_all) && (country_africa || country_all || country_eg || country_northern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEgypt.RegisterName(xlanguage.Korean, "이집트")
	dataEgypt.RegisterOfficialName(xlanguage.Korean, "이집트 아랍 공화국")
	dataEgypt.RegisterCapital(xlanguage.Korean, "카이로")
}

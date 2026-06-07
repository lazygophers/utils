//go:build (lang_ko || lang_all) && (country_africa || country_all || country_lr || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLiberia.RegisterName(xlanguage.Korean, "라이베리아")
	dataLiberia.RegisterOfficialName(xlanguage.Korean, "라이베리아 공화국")
	dataLiberia.RegisterCapital(xlanguage.Korean, "몬로비아")
}

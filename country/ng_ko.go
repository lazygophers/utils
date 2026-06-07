//go:build (lang_ko || lang_all) && (country_africa || country_all || country_ng || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNigeria.RegisterName(xlanguage.Korean, "나이지리아")
	dataNigeria.RegisterOfficialName(xlanguage.Korean, "나이지리아 연방 공화국")
	dataNigeria.RegisterCapital(xlanguage.Korean, "아부자")
}

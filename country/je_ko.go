//go:build (lang_ko || lang_all) && (country_all || country_europe || country_je || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJersey.RegisterName(xlanguage.Korean, "저지")
	dataJersey.RegisterOfficialName(xlanguage.Korean, "저지 구역")
	dataJersey.RegisterCapital(xlanguage.Korean, "세인트헬리어")
}

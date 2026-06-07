//go:build (lang_ko || lang_all) && (country_all || country_europe || country_gg || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuernsey.RegisterName(xlanguage.Korean, "건지")
	dataGuernsey.RegisterOfficialName(xlanguage.Korean, "건지 구역")
	dataGuernsey.RegisterCapital(xlanguage.Korean, "세인트피터포트")
}

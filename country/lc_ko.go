//go:build (lang_ko || lang_all) && (country_all || country_americas || country_caribbean || country_lc)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintLucia.RegisterName(xlanguage.Korean, "세인트루시아")
	dataSaintLucia.RegisterOfficialName(xlanguage.Korean, "세인트루시아")
	dataSaintLucia.RegisterCapital(xlanguage.Korean, "캐스트리스")
}

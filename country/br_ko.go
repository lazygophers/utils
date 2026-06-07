//go:build (lang_ko || lang_all) && (country_all || country_americas || country_br || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrazil.RegisterName(xlanguage.Korean, "브라질")
	dataBrazil.RegisterOfficialName(xlanguage.Korean, "브라질 연방 공화국")
	dataBrazil.RegisterCapital(xlanguage.Korean, "브라질리아")
}

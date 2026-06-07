//go:build (lang_ko || lang_all) && (country_all || country_americas || country_caribbean || country_cu)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCuba.RegisterName(xlanguage.Korean, "쿠바")
	dataCuba.RegisterOfficialName(xlanguage.Korean, "쿠바 공화국")
	dataCuba.RegisterCapital(xlanguage.Korean, "아바나")
}

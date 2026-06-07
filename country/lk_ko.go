//go:build (lang_ko || lang_all) && (country_all || country_asia || country_lk || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSriLanka.RegisterName(xlanguage.Korean, "스리랑카")
	dataSriLanka.RegisterOfficialName(xlanguage.Korean, "스리랑카 민주 사회주의 공화국")
	dataSriLanka.RegisterCapital(xlanguage.Korean, "스리자야와르데네푸라코테")
}

//go:build (lang_ko || lang_all) && (country_all || country_americas || country_caribbean || country_pr)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPuertoRico.RegisterName(xlanguage.Korean, "푸에르토리코")
	dataPuertoRico.RegisterOfficialName(xlanguage.Korean, "푸에르토리코 자치 연방구")
	dataPuertoRico.RegisterCapital(xlanguage.Korean, "산후안")
}

//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCongo.RegisterName(xlanguage.Korean, "콩고 공화국")
	dataCongo.RegisterOfficialName(xlanguage.Korean, "콩고 공화국")
	dataCongo.RegisterCapital(xlanguage.Korean, "브라자빌")
}

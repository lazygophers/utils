//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalawi.RegisterName(xlanguage.Korean, "말라위")
	dataMalawi.RegisterOfficialName(xlanguage.Korean, "말라위 공화국")
	dataMalawi.RegisterCapital(xlanguage.Korean, "릴롱궤")
}

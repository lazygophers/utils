//go:build (lang_ko || lang_all) && (country_africa || country_all || country_mr || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMauritania.RegisterName(xlanguage.Korean, "모리타니")
	dataMauritania.RegisterOfficialName(xlanguage.Korean, "모리타니 이슬람 공화국")
	dataMauritania.RegisterCapital(xlanguage.Korean, "누악쇼트")
}

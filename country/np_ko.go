//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNepal.RegisterName(xlanguage.Korean, "네팔")
	dataNepal.RegisterOfficialName(xlanguage.Korean, "네팔 연방 민주 공화국")
	dataNepal.RegisterCapital(xlanguage.Korean, "카트만두")
}

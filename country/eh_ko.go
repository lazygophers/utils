//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataWesternSahara.RegisterName(xlanguage.Korean, "서사하라")
	dataWesternSahara.RegisterOfficialName(xlanguage.Korean, "사하라 아랍 민주 공화국")
	dataWesternSahara.RegisterCapital(xlanguage.Korean, "엘아이운")
}

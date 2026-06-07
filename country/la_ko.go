//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLaos.RegisterName(xlanguage.Korean, "라오스")
	dataLaos.RegisterOfficialName(xlanguage.Korean, "라오 인민 민주 공화국")
	dataLaos.RegisterCapital(xlanguage.Korean, "비엔티안")
}

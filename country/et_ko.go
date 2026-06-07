//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEthiopia.RegisterName(xlanguage.Korean, "에티오피아")
	dataEthiopia.RegisterOfficialName(xlanguage.Korean, "에티오피아 연방 민주 공화국")
	dataEthiopia.RegisterCapital(xlanguage.Korean, "아디스아바바")
}

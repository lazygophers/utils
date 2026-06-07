//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMarshallIslands.RegisterName(xlanguage.Korean, "마셜 제도")
	dataMarshallIslands.RegisterOfficialName(xlanguage.Korean, "마셜 제도 공화국")
	dataMarshallIslands.RegisterCapital(xlanguage.Korean, "마주로")
}

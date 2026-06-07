//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLebanon.RegisterName(xlanguage.Korean, "레바논")
	dataLebanon.RegisterOfficialName(xlanguage.Korean, "레바논 공화국")
	dataLebanon.RegisterCapital(xlanguage.Korean, "베이루트")
}

//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalta.RegisterName(xlanguage.Korean, "몰타")
	dataMalta.RegisterOfficialName(xlanguage.Korean, "몰타 공화국")
	dataMalta.RegisterCapital(xlanguage.Korean, "발레타")
}

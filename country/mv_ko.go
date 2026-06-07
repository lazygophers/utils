//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMaldives.RegisterName(xlanguage.Korean, "몰디브")
	dataMaldives.RegisterOfficialName(xlanguage.Korean, "몰디브 공화국")
	dataMaldives.RegisterCapital(xlanguage.Korean, "말레")
}

//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKiribati.RegisterName(xlanguage.Korean, "키리바시")
	dataKiribati.RegisterOfficialName(xlanguage.Korean, "키리바시 공화국")
	dataKiribati.RegisterCapital(xlanguage.Korean, "타라와")
}

//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLesotho.RegisterName(xlanguage.Korean, "레소토")
	dataLesotho.RegisterOfficialName(xlanguage.Korean, "레소토 왕국")
	dataLesotho.RegisterCapital(xlanguage.Korean, "마세루")
}

//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSwitzerland.RegisterName(xlanguage.Korean, "스위스")
	dataSwitzerland.RegisterOfficialName(xlanguage.Korean, "스위스 연방")
	dataSwitzerland.RegisterCapital(xlanguage.Korean, "베른")
}

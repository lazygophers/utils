//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSlovakia.RegisterName(xlanguage.Korean, "슬로바키아")
	dataSlovakia.RegisterOfficialName(xlanguage.Korean, "슬로바키아 공화국")
	dataSlovakia.RegisterCapital(xlanguage.Korean, "브라티슬라바")
}

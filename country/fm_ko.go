//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMicronesia.RegisterName(xlanguage.Korean, "미크로네시아 연방")
	dataMicronesia.RegisterOfficialName(xlanguage.Korean, "미크로네시아 연방")
	dataMicronesia.RegisterCapital(xlanguage.Korean, "팔리키르")
}

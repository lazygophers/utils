//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRussia.RegisterName(xlanguage.Korean, "러시아")
	dataRussia.RegisterOfficialName(xlanguage.Korean, "러시아 연방")
	dataRussia.RegisterCapital(xlanguage.Korean, "모스크바")
}

//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSingapore.RegisterName(xlanguage.Korean, "싱가포르")
	dataSingapore.RegisterOfficialName(xlanguage.Korean, "싱가포르 공화국")
	dataSingapore.RegisterCapital(xlanguage.Korean, "싱가포르")
}

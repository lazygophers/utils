//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSweden.RegisterName(xlanguage.Korean, "스웨덴")
	dataSweden.RegisterOfficialName(xlanguage.Korean, "스웨덴 왕국")
	dataSweden.RegisterCapital(xlanguage.Korean, "스톡홀름")
}

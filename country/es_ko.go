//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSpain.RegisterName(xlanguage.Korean, "스페인")
	dataSpain.RegisterOfficialName(xlanguage.Korean, "스페인 왕국")
	dataSpain.RegisterCapital(xlanguage.Korean, "마드리드")
}

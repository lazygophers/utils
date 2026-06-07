//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataThailand.RegisterName(xlanguage.Korean, "태국")
	dataThailand.RegisterOfficialName(xlanguage.Korean, "타이 왕국")
	dataThailand.RegisterCapital(xlanguage.Korean, "방콕")
}

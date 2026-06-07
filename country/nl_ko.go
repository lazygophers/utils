//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNetherlands.RegisterName(xlanguage.Korean, "네덜란드")
	dataNetherlands.RegisterOfficialName(xlanguage.Korean, "네덜란드 왕국")
	dataNetherlands.RegisterCapital(xlanguage.Korean, "암스테르담")
}

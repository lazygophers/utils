//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataOman.RegisterName(xlanguage.Korean, "오만")
	dataOman.RegisterOfficialName(xlanguage.Korean, "오만 술탄국")
	dataOman.RegisterCapital(xlanguage.Korean, "무스카트")
}

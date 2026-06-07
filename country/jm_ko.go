//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJamaica.RegisterName(xlanguage.Korean, "자메이카")
	dataJamaica.RegisterOfficialName(xlanguage.Korean, "자메이카")
	dataJamaica.RegisterCapital(xlanguage.Korean, "킹스턴")
}

//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGrenada.RegisterName(xlanguage.Korean, "그레나다")
	dataGrenada.RegisterOfficialName(xlanguage.Korean, "그레나다")
	dataGrenada.RegisterCapital(xlanguage.Korean, "세인트조지스")
}

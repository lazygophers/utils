//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorfolkIsland.RegisterName(xlanguage.Korean, "노퍽섬")
	dataNorfolkIsland.RegisterOfficialName(xlanguage.Korean, "노퍽섬")
	dataNorfolkIsland.RegisterCapital(xlanguage.Korean, "킹스턴")
}

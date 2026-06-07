//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSolomonIslands.RegisterName(xlanguage.Korean, "솔로몬 제도")
	dataSolomonIslands.RegisterOfficialName(xlanguage.Korean, "솔로몬 제도")
	dataSolomonIslands.RegisterCapital(xlanguage.Korean, "호니아라")
}

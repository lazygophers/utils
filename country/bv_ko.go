//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBouvetIsland.RegisterName(xlanguage.Korean, "부베섬")
	dataBouvetIsland.RegisterOfficialName(xlanguage.Korean, "부베섬")
}

//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHungary.RegisterName(xlanguage.Korean, "헝가리")
	dataHungary.RegisterOfficialName(xlanguage.Korean, "헝가리")
	dataHungary.RegisterCapital(xlanguage.Korean, "부다페스트")
}

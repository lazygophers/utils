//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSomalia.RegisterName(xlanguage.Korean, "소말리아")
	dataSomalia.RegisterOfficialName(xlanguage.Korean, "소말리아 연방 공화국")
	dataSomalia.RegisterCapital(xlanguage.Korean, "모가디슈")
}

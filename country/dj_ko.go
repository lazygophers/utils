//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDjibouti.RegisterName(xlanguage.Korean, "지부티")
	dataDjibouti.RegisterOfficialName(xlanguage.Korean, "지부티 공화국")
	dataDjibouti.RegisterCapital(xlanguage.Korean, "지부티")
}

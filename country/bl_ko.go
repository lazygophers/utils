//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintBarthelemy.RegisterName(xlanguage.Korean, "생바르텔레미")
	dataSaintBarthelemy.RegisterOfficialName(xlanguage.Korean, "생바르텔레미 집합체")
	dataSaintBarthelemy.RegisterCapital(xlanguage.Korean, "구스타비아")
}

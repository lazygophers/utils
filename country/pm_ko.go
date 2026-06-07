//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintPierreAndMiquelon.RegisterName(xlanguage.Korean, "생피에르 미클롱")
	dataSaintPierreAndMiquelon.RegisterOfficialName(xlanguage.Korean, "생피에르 미클롱 자치 행정구")
	dataSaintPierreAndMiquelon.RegisterCapital(xlanguage.Korean, "생피에르")
}

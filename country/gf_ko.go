//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchGuiana.RegisterName(xlanguage.Korean, "프랑스령 기아나")
	dataFrenchGuiana.RegisterOfficialName(xlanguage.Korean, "프랑스령 기아나")
	dataFrenchGuiana.RegisterCapital(xlanguage.Korean, "카옌")
}

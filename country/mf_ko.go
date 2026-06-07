//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintMartin.RegisterName(xlanguage.Korean, "생마르탱")
	dataSaintMartin.RegisterOfficialName(xlanguage.Korean, "생마르탱 집합체")
	dataSaintMartin.RegisterCapital(xlanguage.Korean, "마리고")
}

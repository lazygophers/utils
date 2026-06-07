//go:build (lang_ko || lang_all) && (country_all || country_americas || country_caribbean || country_do)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDominicanRepublic.RegisterName(xlanguage.Korean, "도미니카 공화국")
	dataDominicanRepublic.RegisterOfficialName(xlanguage.Korean, "도미니카 공화국")
	dataDominicanRepublic.RegisterCapital(xlanguage.Korean, "산토도밍고")
}

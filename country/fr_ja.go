//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrance.RegisterName(xlanguage.Japanese, "フランス")
	dataFrance.RegisterOfficialName(xlanguage.Japanese, "フランス共和国")
	dataFrance.RegisterCapital(xlanguage.Japanese, "パリ")
}

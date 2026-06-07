//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchPolynesia.RegisterName(xlanguage.Japanese, "フランス領ポリネシア")
	dataFrenchPolynesia.RegisterOfficialName(xlanguage.Japanese, "フランス領ポリネシア")
	dataFrenchPolynesia.RegisterCapital(xlanguage.Japanese, "パペーテ")
}

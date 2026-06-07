//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHaiti.RegisterName(xlanguage.Japanese, "ハイチ")
	dataHaiti.RegisterOfficialName(xlanguage.Japanese, "ハイチ共和国")
	dataHaiti.RegisterCapital(xlanguage.Japanese, "ポルトープランス")
}

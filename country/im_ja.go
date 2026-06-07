//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsleOfMan.RegisterName(xlanguage.Japanese, "マン島")
	dataIsleOfMan.RegisterOfficialName(xlanguage.Japanese, "マン島")
	dataIsleOfMan.RegisterCapital(xlanguage.Japanese, "ダグラス")
}

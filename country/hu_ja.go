//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHungary.RegisterName(xlanguage.Japanese, "ハンガリー")
	dataHungary.RegisterOfficialName(xlanguage.Japanese, "ハンガリー")
	dataHungary.RegisterCapital(xlanguage.Japanese, "ブダペスト")
}

//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuernsey.RegisterName(xlanguage.Japanese, "ガーンジー")
	dataGuernsey.RegisterOfficialName(xlanguage.Japanese, "ガーンジー")
	dataGuernsey.RegisterCapital(xlanguage.Japanese, "セント・ピーター・ポート")
}

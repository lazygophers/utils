//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJersey.RegisterName(xlanguage.Japanese, "ジャージー")
	dataJersey.RegisterOfficialName(xlanguage.Japanese, "ジャージー")
	dataJersey.RegisterCapital(xlanguage.Japanese, "セント・ヘリア")
}

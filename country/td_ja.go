//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChad.RegisterName(xlanguage.Japanese, "チャド")
	dataChad.RegisterOfficialName(xlanguage.Japanese, "チャド共和国")
	dataChad.RegisterCapital(xlanguage.Japanese, "ンジャメナ")
}

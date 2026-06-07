//go:build lang_ja || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bif.RegisterName(xlanguage.Japanese, "ブルンジ・フラン")
}

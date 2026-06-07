//go:build lang_ja || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mop.RegisterName(xlanguage.Japanese, "マカオ・パタカ")
}

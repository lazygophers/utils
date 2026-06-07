//go:build lang_ja || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Zmw.RegisterName(xlanguage.Japanese, "ザンビア・クワチャ")
}

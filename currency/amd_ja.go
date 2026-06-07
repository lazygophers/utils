//go:build lang_ja || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Amd.RegisterName(xlanguage.Japanese, "アルメニア・ドラム")
}

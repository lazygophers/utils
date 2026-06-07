//go:build lang_ja || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Lyd.RegisterName(xlanguage.Japanese, "リビア・ディナール")
}

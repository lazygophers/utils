//go:build lang_ja || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mnt.RegisterName(xlanguage.Japanese, "トゥグルグ")
}

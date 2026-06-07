//go:build lang_ja || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Uah.RegisterName(xlanguage.Japanese, "フリヴニャ")
}

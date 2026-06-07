//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Lkr.RegisterName(xlanguage.MustParse("zh-Hant"), "斯里蘭卡盧比")
}

//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Btn.RegisterName(xlanguage.MustParse("zh-Hant"), "不丹努爾特魯姆")
}

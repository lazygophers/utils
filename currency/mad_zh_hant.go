//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mad.RegisterName(xlanguage.MustParse("zh-Hant"), "摩洛哥迪拉姆")
}

//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Nio.RegisterName(xlanguage.MustParse("zh-Hant"), "尼加拉瓜科多巴")
}

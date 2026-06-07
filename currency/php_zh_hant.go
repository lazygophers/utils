//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Php.RegisterName(xlanguage.MustParse("zh-Hant"), "菲律賓披索")
}

//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSeychelles.RegisterName(xlanguage.MustParse("zh-Hant"), "塞席爾")
	dataSeychelles.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "塞席爾共和國")
	dataSeychelles.RegisterCapital(xlanguage.MustParse("zh-Hant"), "維多利亞")
}

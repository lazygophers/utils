//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMadagascar.RegisterName(xlanguage.MustParse("zh-Hant"), "馬達加斯加")
	dataMadagascar.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "馬達加斯加共和國")
	dataMadagascar.RegisterCapital(xlanguage.MustParse("zh-Hant"), "安塔那那利佛")
}

//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBahamas.RegisterName(xlanguage.MustParse("zh-Hant"), "巴哈馬")
	dataBahamas.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "巴哈馬國")
	dataBahamas.RegisterCapital(xlanguage.MustParse("zh-Hant"), "拿索")
}

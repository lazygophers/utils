//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalta.RegisterName(xlanguage.MustParse("zh-Hant"), "馬爾他")
	dataMalta.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "馬爾他共和國")
	dataMalta.RegisterCapital(xlanguage.MustParse("zh-Hant"), "瓦萊塔")
}

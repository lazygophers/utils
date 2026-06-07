//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaudiArabia.RegisterName(xlanguage.MustParse("zh-Hant"), "沙烏地阿拉伯")
	dataSaudiArabia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "沙烏地阿拉伯王國")
	dataSaudiArabia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "利雅德")
}

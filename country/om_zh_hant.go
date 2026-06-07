//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataOman.RegisterName(xlanguage.MustParse("zh-Hant"), "阿曼")
	dataOman.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "阿曼蘇丹國")
	dataOman.RegisterCapital(xlanguage.MustParse("zh-Hant"), "馬斯喀特")
}

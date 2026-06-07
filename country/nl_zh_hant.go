//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNetherlands.RegisterName(xlanguage.MustParse("zh-Hant"), "荷蘭")
	dataNetherlands.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "荷蘭王國")
	dataNetherlands.RegisterCapital(xlanguage.MustParse("zh-Hant"), "阿姆斯特丹")
}

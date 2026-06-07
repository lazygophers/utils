//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPortugal.RegisterName(xlanguage.MustParse("zh-Hant"), "葡萄牙")
	dataPortugal.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "葡萄牙共和國")
	dataPortugal.RegisterCapital(xlanguage.MustParse("zh-Hant"), "里斯本")
}

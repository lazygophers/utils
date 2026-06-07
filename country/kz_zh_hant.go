//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKazakhstan.RegisterName(xlanguage.MustParse("zh-Hant"), "哈薩克")
	dataKazakhstan.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "哈薩克共和國")
	dataKazakhstan.RegisterCapital(xlanguage.MustParse("zh-Hant"), "阿斯塔納")
}

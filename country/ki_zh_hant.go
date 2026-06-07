//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKiribati.RegisterName(xlanguage.MustParse("zh-Hant"), "吉里巴斯")
	dataKiribati.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "吉里巴斯共和國")
	dataKiribati.RegisterCapital(xlanguage.MustParse("zh-Hant"), "塔拉瓦")
}

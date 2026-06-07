//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataWesternSahara.RegisterName(xlanguage.MustParse("zh-Hant"), "西撒哈拉")
	dataWesternSahara.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "阿拉伯撒哈拉民主共和國")
	dataWesternSahara.RegisterCapital(xlanguage.MustParse("zh-Hant"), "阿尤恩")
}

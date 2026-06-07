//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFiji.RegisterName(xlanguage.MustParse("zh-Hant"), "斐濟")
	dataFiji.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "斐濟共和國")
	dataFiji.RegisterCapital(xlanguage.MustParse("zh-Hant"), "蘇瓦")
}

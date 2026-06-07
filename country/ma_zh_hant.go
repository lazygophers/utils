//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMorocco.RegisterName(xlanguage.MustParse("zh-Hant"), "摩洛哥")
	dataMorocco.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "摩洛哥王國")
	dataMorocco.RegisterCapital(xlanguage.MustParse("zh-Hant"), "拉巴特")
}

//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMonaco.RegisterName(xlanguage.MustParse("zh-Hant"), "摩納哥")
	dataMonaco.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "摩納哥親王國")
	dataMonaco.RegisterCapital(xlanguage.MustParse("zh-Hant"), "摩納哥")
}

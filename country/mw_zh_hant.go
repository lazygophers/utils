//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalawi.RegisterName(xlanguage.MustParse("zh-Hant"), "馬拉威")
	dataMalawi.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "馬拉威共和國")
	dataMalawi.RegisterCapital(xlanguage.MustParse("zh-Hant"), "里朗威")
}

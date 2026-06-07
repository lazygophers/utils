//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZimbabwe.RegisterName(xlanguage.MustParse("zh-Hant"), "辛巴威")
	dataZimbabwe.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "辛巴威共和國")
	dataZimbabwe.RegisterCapital(xlanguage.MustParse("zh-Hant"), "哈拉雷")
}

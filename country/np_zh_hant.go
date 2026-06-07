//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNepal.RegisterName(xlanguage.MustParse("zh-Hant"), "尼泊爾")
	dataNepal.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "尼泊爾聯邦民主共和國")
	dataNepal.RegisterCapital(xlanguage.MustParse("zh-Hant"), "加德滿都")
}

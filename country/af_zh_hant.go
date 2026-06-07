//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAfghanistan.RegisterName(xlanguage.MustParse("zh-Hant"), "阿富汗")
	dataAfghanistan.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "阿富汗伊斯蘭酋長國")
	dataAfghanistan.RegisterCapital(xlanguage.MustParse("zh-Hant"), "喀布爾")
}

//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_pk || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPakistan.RegisterName(xlanguage.MustParse("zh-Hant"), "巴基斯坦")
	dataPakistan.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "巴基斯坦伊斯蘭共和國")
	dataPakistan.RegisterCapital(xlanguage.MustParse("zh-Hant"), "伊斯蘭瑪巴德")
}

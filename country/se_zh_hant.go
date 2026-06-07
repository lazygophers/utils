//go:build (lang_zh_hant || lang_all) && (country_all || country_europe || country_northern_europe || country_se)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSweden.RegisterName(xlanguage.MustParse("zh-Hant"), "瑞典")
	dataSweden.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "瑞典王國")
	dataSweden.RegisterCapital(xlanguage.MustParse("zh-Hant"), "斯德哥爾摩")
}

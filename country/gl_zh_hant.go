//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_gl || country_northern_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGreenland.RegisterName(xlanguage.MustParse("zh-Hant"), "格陵蘭")
	dataGreenland.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "格陵蘭")
	dataGreenland.RegisterCapital(xlanguage.MustParse("zh-Hant"), "努克")
}

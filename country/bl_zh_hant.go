//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_bl || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintBarthelemy.RegisterName(xlanguage.MustParse("zh-Hant"), "聖巴泰勒米")
	dataSaintBarthelemy.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "聖巴泰勒米集體")
	dataSaintBarthelemy.RegisterCapital(xlanguage.MustParse("zh-Hant"), "古斯塔維亞")
}

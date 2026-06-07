//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_central_asia || country_tm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkmenistan.RegisterName(xlanguage.MustParse("zh-Hant"), "土庫曼")
	dataTurkmenistan.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "土庫曼")
	dataTurkmenistan.RegisterCapital(xlanguage.MustParse("zh-Hant"), "阿什哈巴特")
}

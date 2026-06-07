//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_lk || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSriLanka.RegisterName(xlanguage.MustParse("zh-Hant"), "斯里蘭卡")
	dataSriLanka.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "斯里蘭卡民主社會主義共和國")
	dataSriLanka.RegisterCapital(xlanguage.MustParse("zh-Hant"), "斯里賈亞瓦德納普拉科特")
}

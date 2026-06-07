//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_bt || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBhutan.RegisterName(xlanguage.MustParse("zh-Hant"), "不丹")
	dataBhutan.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "不丹王國")
	dataBhutan.RegisterCapital(xlanguage.MustParse("zh-Hant"), "辛布")
}

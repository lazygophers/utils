//go:build (lang_zh_hant || lang_all) && (country_all || country_ba || country_europe || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBosniaAndHerzegovina.RegisterName(xlanguage.MustParse("zh-Hant"), "波士尼亞與赫塞哥維納")
	dataBosniaAndHerzegovina.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "波士尼亞與赫塞哥維納")
	dataBosniaAndHerzegovina.RegisterCapital(xlanguage.MustParse("zh-Hant"), "塞拉耶佛")
}

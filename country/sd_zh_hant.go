//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_northern_africa || country_sd)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSudan.RegisterName(xlanguage.MustParse("zh-Hant"), "蘇丹")
	dataSudan.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "蘇丹共和國")
	dataSudan.RegisterCapital(xlanguage.MustParse("zh-Hant"), "喀土穆")
}

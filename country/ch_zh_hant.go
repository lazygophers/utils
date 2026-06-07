//go:build (lang_zh_hant || lang_all) && (country_all || country_ch || country_europe || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSwitzerland.RegisterName(xlanguage.MustParse("zh-Hant"), "瑞士")
	dataSwitzerland.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "瑞士聯邦")
	dataSwitzerland.RegisterCapital(xlanguage.MustParse("zh-Hant"), "伯恩")
}

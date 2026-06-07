//go:build (lang_zh_hant || lang_all) && (country_all || country_fm || country_micronesia || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMicronesia.RegisterName(xlanguage.MustParse("zh-Hant"), "密克羅尼西亞聯邦")
	dataMicronesia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "密克羅尼西亞聯邦")
	dataMicronesia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "帕利基爾")
}

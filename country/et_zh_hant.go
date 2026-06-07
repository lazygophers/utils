//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_eastern_africa || country_et)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEthiopia.RegisterName(xlanguage.MustParse("zh-Hant"), "衣索比亞")
	dataEthiopia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "衣索比亞聯邦民主共和國")
	dataEthiopia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "阿迪斯阿貝巴")
}

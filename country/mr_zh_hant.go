//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_mr || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMauritania.RegisterName(xlanguage.MustParse("zh-Hant"), "茅利塔尼亞")
	dataMauritania.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "茅利塔尼亞伊斯蘭共和國")
	dataMauritania.RegisterCapital(xlanguage.MustParse("zh-Hant"), "諾克少")
}

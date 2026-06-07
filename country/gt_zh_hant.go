//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_central_america || country_gt)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuatemala.RegisterName(xlanguage.MustParse("zh-Hant"), "瓜地馬拉")
	dataGuatemala.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "瓜地馬拉共和國")
	dataGuatemala.RegisterCapital(xlanguage.MustParse("zh-Hant"), "瓜地馬拉市")
}

//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_bb || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBarbados.RegisterName(xlanguage.MustParse("zh-Hant"), "巴貝多")
	dataBarbados.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "巴貝多")
	dataBarbados.RegisterCapital(xlanguage.MustParse("zh-Hant"), "橋鎮")
}

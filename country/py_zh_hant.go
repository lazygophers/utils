//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_py || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataParaguay.RegisterName(xlanguage.MustParse("zh-Hant"), "巴拉圭")
	dataParaguay.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "巴拉圭共和國")
	dataParaguay.RegisterCapital(xlanguage.MustParse("zh-Hant"), "亞松森")
}

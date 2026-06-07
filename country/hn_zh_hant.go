//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_central_america || country_hn)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHonduras.RegisterName(xlanguage.MustParse("zh-Hant"), "宏都拉斯")
	dataHonduras.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "宏都拉斯共和國")
	dataHonduras.RegisterCapital(xlanguage.MustParse("zh-Hant"), "德古斯加巴")
}

//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_br || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrazil.RegisterName(xlanguage.MustParse("zh-Hant"), "巴西")
	dataBrazil.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "巴西聯邦共和國")
	dataBrazil.RegisterCapital(xlanguage.MustParse("zh-Hant"), "巴西利亞")
}

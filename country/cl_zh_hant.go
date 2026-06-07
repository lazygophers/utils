//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_cl || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChile.RegisterName(xlanguage.MustParse("zh-Hant"), "智利")
	dataChile.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "智利共和國")
	dataChile.RegisterCapital(xlanguage.MustParse("zh-Hant"), "聖地牙哥")
}

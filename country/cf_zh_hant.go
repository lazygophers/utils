//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_cf || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCentralAfricanRepublic.RegisterName(xlanguage.MustParse("zh-Hant"), "中非共和國")
	dataCentralAfricanRepublic.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "中非共和國")
	dataCentralAfricanRepublic.RegisterCapital(xlanguage.MustParse("zh-Hant"), "班基")
}

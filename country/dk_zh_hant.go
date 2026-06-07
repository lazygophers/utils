//go:build (lang_zh_hant || lang_all) && (country_all || country_dk || country_europe || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDenmark.RegisterName(xlanguage.MustParse("zh-Hant"), "丹麥")
	dataDenmark.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "丹麥王國")
	dataDenmark.RegisterCapital(xlanguage.MustParse("zh-Hant"), "哥本哈根")
}

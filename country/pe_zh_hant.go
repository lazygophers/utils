//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_pe || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPeru.RegisterName(xlanguage.MustParse("zh-Hant"), "秘魯")
	dataPeru.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "秘魯共和國")
	dataPeru.RegisterCapital(xlanguage.MustParse("zh-Hant"), "利馬")
}

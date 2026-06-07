//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_mv || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMaldives.RegisterName(xlanguage.MustParse("zh-Hant"), "馬爾地夫")
	dataMaldives.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "馬爾地夫共和國")
	dataMaldives.RegisterCapital(xlanguage.MustParse("zh-Hant"), "馬列")
}

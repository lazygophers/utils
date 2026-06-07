//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_bw || country_southern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBotswana.RegisterName(xlanguage.MustParse("zh-Hant"), "波札那")
	dataBotswana.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "波札那共和國")
	dataBotswana.RegisterCapital(xlanguage.MustParse("zh-Hant"), "嘉伯隆里")
}

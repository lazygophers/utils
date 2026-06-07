//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_eastern_asia || country_kp)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthKorea.RegisterName(xlanguage.MustParse("zh-Hant"), "北韓")
	dataNorthKorea.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "朝鮮民主主義人民共和國")
	dataNorthKorea.RegisterCapital(xlanguage.MustParse("zh-Hant"), "平壤")
}

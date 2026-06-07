//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_ne || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNiger.RegisterName(xlanguage.MustParse("zh-Hant"), "尼日")
	dataNiger.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "尼日共和國")
	dataNiger.RegisterCapital(xlanguage.MustParse("zh-Hant"), "尼阿美")
}

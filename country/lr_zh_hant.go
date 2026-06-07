//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_lr || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLiberia.RegisterName(xlanguage.MustParse("zh-Hant"), "賴比瑞亞")
	dataLiberia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "賴比瑞亞共和國")
	dataLiberia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "蒙羅維亞")
}

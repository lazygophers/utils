//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_central_america || country_mx)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMexico.RegisterName(xlanguage.MustParse("zh-Hant"), "墨西哥")
	dataMexico.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "墨西哥合眾國")
	dataMexico.RegisterCapital(xlanguage.MustParse("zh-Hant"), "墨西哥城")
}

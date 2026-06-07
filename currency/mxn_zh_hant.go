//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_central_america || country_mx || currency_all || currency_mxn)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mxn.RegisterName(xlanguage.MustParse("zh-Hant"), "墨西哥披索")
}

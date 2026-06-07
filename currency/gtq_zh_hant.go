//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_central_america || country_gt || currency_all || currency_gtq)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Gtq.RegisterName(xlanguage.MustParse("zh-Hant"), "格查爾")
}

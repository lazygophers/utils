//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_cv || country_western_africa || currency_all || currency_cve)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	CVE.RegisterName(xlanguage.MustParse("zh-Hant"), "維德角埃斯庫多")
}

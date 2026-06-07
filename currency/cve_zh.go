//go:build country_africa || country_all || country_cv || country_western_africa || currency_all || currency_cve

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	CVE.RegisterName(xlanguage.Chinese, "佛得角埃斯库多")
}

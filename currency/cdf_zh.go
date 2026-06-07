//go:build country_africa || country_all || country_cd || country_middle_africa || currency_all || currency_cdf

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Cdf.RegisterName(xlanguage.Chinese, "刚果法郎")
}

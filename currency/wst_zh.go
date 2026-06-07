//go:build country_all || country_oceania || country_polynesia || country_ws || currency_all || currency_wst

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Wst.RegisterName(xlanguage.Chinese, "萨摩亚塔拉")
}

//go:build country_africa || country_all || country_bf || country_bj || country_ci || country_gw || country_ml || country_ne || country_sn || country_tg || country_western_africa || currency_all || currency_xof

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Xof.RegisterName(xlanguage.Chinese, "西非法郎")
}

//go:build (lang_es || lang_all) && (country_africa || country_all || country_bf || country_bj || country_ci || country_gw || country_ml || country_ne || country_sn || country_tg || country_western_africa || currency_all || currency_xof)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	XOF.RegisterName(xlanguage.Spanish, "Franco CFA de África Occidental")
}

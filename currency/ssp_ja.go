//go:build (lang_ja || lang_all) && (country_africa || country_all || country_eastern_africa || country_ss || currency_all || currency_ssp)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	SSP.RegisterName(xlanguage.Japanese, "南スーダン・ポンド")
}

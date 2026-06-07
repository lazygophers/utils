//go:build country_africa || country_all || country_eastern_africa || country_mw || currency_all || currency_mwk

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	MWK.RegisterName(xlanguage.Chinese, "马拉维克瓦查")
}

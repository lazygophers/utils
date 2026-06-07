//go:build country_africa || country_all || country_eastern_africa || country_et || currency_all || currency_etb

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Etb.RegisterName(xlanguage.Chinese, "埃塞俄比亚比尔")
}

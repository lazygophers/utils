//go:build country_africa || country_all || country_eastern_africa || country_er || currency_all || currency_ern

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Ern.RegisterName(xlanguage.Chinese, "厄立特里亚纳克法")
}

//go:build country_all || country_europe || country_northern_europe || country_se || currency_all || currency_sek

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	SEK.RegisterName(xlanguage.Chinese, "瑞典克朗")
}

//go:build country_africa || country_all || country_na || country_southern_africa || currency_all || currency_nad

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Nad.RegisterName(xlanguage.Chinese, "纳米比亚元")
}

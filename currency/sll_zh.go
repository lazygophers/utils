//go:build country_africa || country_all || country_sl || country_western_africa || currency_all || currency_sll

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	SLL.RegisterName(xlanguage.Chinese, "塞拉利昂利昂")
}

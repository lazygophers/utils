//go:build country_africa || country_all || country_eastern_africa || country_sc || currency_all || currency_scr

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	SCR.RegisterName(xlanguage.Chinese, "塞舌尔卢比")
}

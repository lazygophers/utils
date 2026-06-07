//go:build (lang_ko || lang_all) && (country_all || country_americas || country_aw || country_caribbean || currency_all || currency_awg)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	AWG.RegisterName(xlanguage.Korean, "아루바 플로린")
}

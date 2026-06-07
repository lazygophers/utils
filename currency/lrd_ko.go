//go:build (lang_ko || lang_all) && (country_africa || country_all || country_lr || country_western_africa || currency_all || currency_lrd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	LRD.RegisterName(xlanguage.Korean, "라이베리아 달러")
}

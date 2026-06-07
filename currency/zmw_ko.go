//go:build (lang_ko || lang_all) && (country_africa || country_all || country_eastern_africa || country_zm || currency_all || currency_zmw)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	ZMW.RegisterName(xlanguage.Korean, "잠비아 콰차")
}

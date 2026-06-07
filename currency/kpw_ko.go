//go:build (lang_ko || lang_all) && (country_all || country_asia || country_eastern_asia || country_kp || currency_all || currency_kpw)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Kpw.RegisterName(xlanguage.Korean, "조선민주주의인민공화국 원")
}

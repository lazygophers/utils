//go:build (lang_ko || lang_all) && (country_all || country_americas || country_caribbean || country_tt || currency_all || currency_ttd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	TTD.RegisterName(xlanguage.Korean, "트리니다드 토바고 달러")
}

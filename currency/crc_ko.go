//go:build (lang_ko || lang_all) && (country_all || country_americas || country_central_america || country_cr || currency_all || currency_crc)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	CRC.RegisterName(xlanguage.Korean, "코스타리카 콜론")
}

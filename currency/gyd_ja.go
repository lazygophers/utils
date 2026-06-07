//go:build (lang_ja || lang_all) && (country_all || country_americas || country_gy || country_south_america || currency_all || currency_gyd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	GYD.RegisterName(xlanguage.Japanese, "ガイアナ・ドル")
}

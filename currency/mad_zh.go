//go:build country_africa || country_all || country_eh || country_ma || country_northern_africa || currency_all || currency_mad

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	MAD.RegisterName(xlanguage.Chinese, "摩洛哥迪拉姆")
}

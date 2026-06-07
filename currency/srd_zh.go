//go:build country_all || country_americas || country_south_america || country_sr || currency_all || currency_srd

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Srd.RegisterName(xlanguage.Chinese, "苏里南元")
}

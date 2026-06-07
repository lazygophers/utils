//go:build country_all || country_asia || country_la || country_south_eastern_asia || currency_all || currency_lak

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	LAK.RegisterName(xlanguage.Chinese, "老挝基普")
}

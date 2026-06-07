//go:build country_all || country_asia || country_kh || country_south_eastern_asia || currency_all || currency_khr

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Khr.RegisterName(xlanguage.Chinese, "柬埔寨瑞尔")
}

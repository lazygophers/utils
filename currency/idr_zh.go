//go:build country_all || country_asia || country_id || country_south_eastern_asia || currency_all || currency_idr

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	IDR.RegisterName(xlanguage.Chinese, "印度尼西亚盾")
}

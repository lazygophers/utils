//go:build country_all || country_am || country_asia || country_western_asia || currency_all || currency_amd

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	AMD.RegisterName(xlanguage.Chinese, "亚美尼亚德拉姆")
}

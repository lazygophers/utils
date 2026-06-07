//go:build country_all || country_asia || country_lb || country_western_asia || currency_all || currency_lbp

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Lbp.RegisterName(xlanguage.Chinese, "黎巴嫩镑")
}

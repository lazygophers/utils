//go:build country_all || country_asia || country_central_asia || country_tm || currency_all || currency_tmt

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Tmt.RegisterName(xlanguage.Chinese, "土库曼斯坦马纳特")
}

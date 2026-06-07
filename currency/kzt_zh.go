//go:build country_all || country_asia || country_central_asia || country_kz || currency_all || currency_kzt

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	KZT.RegisterName(xlanguage.Chinese, "哈萨克斯坦坚戈")
}

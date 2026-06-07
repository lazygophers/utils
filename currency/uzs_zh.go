//go:build country_all || country_asia || country_central_asia || country_uz || currency_all || currency_uzs

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	UZS.RegisterName(xlanguage.Chinese, "乌兹别克斯坦苏姆")
}

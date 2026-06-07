//go:build country_all || country_americas || country_cl || country_south_america || currency_all || currency_clp

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Clp.RegisterName(xlanguage.Chinese, "智利比索")
}

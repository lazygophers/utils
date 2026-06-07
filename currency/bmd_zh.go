//go:build country_all || country_americas || country_bm || country_northern_america || currency_all || currency_bmd

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BMD.RegisterName(xlanguage.Chinese, "百慕大元")
}

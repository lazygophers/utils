//go:build country_all || country_europe || country_is || country_northern_europe || currency_all || currency_isk

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Isk.RegisterName(xlanguage.Chinese, "冰岛克朗")
}

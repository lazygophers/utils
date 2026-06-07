//go:build country_all || country_americas || country_bo || country_south_america || currency_all || currency_bob

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bob.RegisterName(xlanguage.Chinese, "玻利维亚诺")
}

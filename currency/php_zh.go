//go:build country_all || country_asia || country_ph || country_south_eastern_asia || currency_all || currency_php

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Php.RegisterName(xlanguage.Chinese, "菲律宾比索")
}

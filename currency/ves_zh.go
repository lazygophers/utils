//go:build country_all || country_americas || country_south_america || country_ve || currency_all || currency_ves

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Ves.RegisterName(xlanguage.Chinese, "委内瑞拉玻利瓦尔")
}

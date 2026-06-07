//go:build country_all || country_americas || country_central_america || country_cr || currency_all || currency_crc

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Crc.RegisterName(xlanguage.Chinese, "哥斯达黎加科朗")
}

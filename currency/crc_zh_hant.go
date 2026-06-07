//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_central_america || country_cr || currency_all || currency_crc)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	CRC.RegisterName(xlanguage.MustParse("zh-Hant"), "哥斯大黎加科朗")
}
